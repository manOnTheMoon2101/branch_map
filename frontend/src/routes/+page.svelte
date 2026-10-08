<script lang="ts">
  import {
    Map,
    MapMarker,
    MarkerContent,
    MarkerPopup,
    MapControls,
  } from "#lib/components/ui/map/index.js";
  import Building2 from "@lucide/svelte/icons/building-2";
  import Landmark from "@lucide/svelte/icons/landmark";
  import MapPin from "@lucide/svelte/icons/map-pin";
  import Phone from "@lucide/svelte/icons/phone";
  import Clock from "@lucide/svelte/icons/clock";
  import ExternalLink from "@lucide/svelte/icons/external-link";
  import CircleAlert from "@lucide/svelte/icons/circle-alert";
  import type { Map as MaplibreMap } from "maplibre-gl";

  type Branch = {
    id: number;
    name: string;
    type: string;
    city: string;
    province: string;
    phone: string;
    hours: string;
    latitude: number;
    longitude: number;
  };

  let branches = $state<Branch[]>([]);
  let fetchError = $state<string | null>(null);
  let mapInstance = $state<MaplibreMap | null>(null);

  // Log coordinates on every map click
  $effect(() => {
    if (!mapInstance) return;

    const handleClick = (e: { lngLat: { lng: number; lat: number } }) => {
      const { lng, lat } = e.lngLat;
      console.log(`Map click → lng: ${lng.toFixed(6)}, lat: ${lat.toFixed(6)}`);
    };

    mapInstance.on("click", handleClick);
    return () => mapInstance?.off("click", handleClick);
  });

  $effect(() => {
    fetch("http://localhost:8080/api/branches")
      .then((r) => {
        if (!r.ok) throw new Error(`Server responded with ${r.status}`);
        return r.json() as Promise<Branch[]>;
      })
      .then((data) => {
        branches = data;
      })
      .catch((err: Error) => {
        fetchError = err.message;
      });
  });
</script>

<div class="relative h-screen w-full overflow-hidden">
  <Map center={[25.0, -29.0]} zoom={5} bind:map={mapInstance}>
    <MapControls showZoom />

    {#each branches as branch (branch.id)}
      <MapMarker longitude={branch.longitude} latitude={branch.latitude}>
        <MarkerContent>
          {#if branch.type === "head"}
            <div
              class="flex h-8 w-8 items-center justify-center rounded-full bg-amber-500 text-white shadow-lg ring-2 ring-white dark:ring-zinc-900"
            >
              <Landmark class="h-3.5 w-3.5" />
            </div>
          {:else}
            <div
              class="flex h-9 w-9 items-center justify-center rounded-full bg-sky-600 text-white shadow-lg ring-2 ring-white dark:ring-zinc-900"
            >
              <Building2 class="h-4 w-4" />
            </div>
          {/if}
        </MarkerContent>

        <MarkerPopup closeButton>
          <div class="w-56 space-y-3">
            <div>
              <h3 class="pr-4 text-sm font-semibold leading-snug">
                {branch.name}
              </h3>
              <span
                class={[
                  "mt-1 inline-block rounded-full px-2 py-0.5 text-xs font-medium capitalize",
                  branch.type === "atm"
                    ? "bg-amber-100 text-amber-700 dark:bg-amber-900/40 dark:text-amber-300"
                    : "bg-sky-100 text-sky-700 dark:bg-sky-900/40 dark:text-sky-300",
                ].join(" ")}
              >
                {branch.type}
              </span>
            </div>

            <div class="space-y-2 text-xs text-muted-foreground">
              <div class="flex items-start gap-2">
                <MapPin
                  class={[
                    "mt-0.5 h-3.5 w-3.5 shrink-0",
                    branch.type === "atm"
                      ? "text-amber-500"
                      : "text-sky-600 dark:text-sky-400",
                  ].join(" ")}
                />
                <span>{branch.city}, {branch.province}</span>
              </div>

              {#if branch.phone}
                <div class="flex items-center gap-2">
                  <Phone
                    class={[
                      "h-3.5 w-3.5 shrink-0",
                      branch.type === "atm"
                        ? "text-amber-500"
                        : "text-sky-600 dark:text-sky-400",
                    ].join(" ")}
                  />
                  <span>{branch.phone}</span>
                </div>
              {/if}

              <div class="flex items-start gap-2">
                <Clock
                  class={[
                    "mt-0.5 h-3.5 w-3.5 shrink-0",
                    branch.type === "atm"
                      ? "text-amber-500"
                      : "text-sky-600 dark:text-sky-400",
                  ].join(" ")}
                />
                <span>{branch.hours}</span>
              </div>
            </div>
          </div>
        </MarkerPopup>
      </MapMarker>
    {/each}
  </Map>

  {#if fetchError}
    <div
      class="absolute bottom-6 left-1/2 flex -translate-x-1/2 items-center gap-2 rounded-lg border bg-background px-4 py-2.5 text-sm text-destructive shadow-lg"
    >
      <CircleAlert class="h-4 w-4 shrink-0" />
      <span>Could not load branches — {fetchError}</span>
    </div>
  {/if}
</div>
