<script lang="ts">
  import {
    Map,
    MapMarker,
    MarkerContent,
    MarkerPopup,
    MapControls,
  } from "#lib/components/ui/map/index.js";
  import Landmark from "@lucide/svelte/icons/landmark";
  import MapPin from "@lucide/svelte/icons/map-pin";
  import Phone from "@lucide/svelte/icons/phone";
  import Clock from "@lucide/svelte/icons/clock";
  import * as Avatar from "#lib/components/ui/avatar/index.js";
  import capitecLogo from "../assets/capitec.png";
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
  <Map
    center={[18.83271, -33.964493]}
    zoom={10}
    bind:map={mapInstance}
    styles={{
      light: "https://basemaps.cartocdn.com/gl/voyager-gl-style/style.json",
      dark: "https://basemaps.cartocdn.com/gl/voyager-gl-style/style.json",
    }}
  >
    <MapControls showZoom showLocate />

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
              class="flex h-9 w-9 items-center justify-center rounded-full bg-white text-white shadow-lg"
            >
              <Avatar.Root>
                <Avatar.Image src={capitecLogo} alt="Capitec Bank" />
                <Avatar.Fallback>CB</Avatar.Fallback>
              </Avatar.Root>
            </div>
          {/if}
        </MarkerContent>

        <MarkerPopup closeButton>
          <div class="w-56 space-y-3">
            <div>
              <h3 class="pr-4 text-sm font-semibold leading-snug">
                {branch.name}
              </h3>
              {#if branch.type == "head"}
                <span>Head Office</span>
              {/if}
            </div>

            <div class="space-y-2 text-xs text-muted-foreground">
              <div class="flex items-start gap-2">
                <MapPin
                  class={[
                    "mt-0.5 h-3.5 w-3.5 shrink-0",
                    branch.type === "head"
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
                      branch.type === "head"
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
                    branch.type === "head"
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
