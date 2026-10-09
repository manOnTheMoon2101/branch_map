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
  import ChevronDown from "@lucide/svelte/icons/chevron-down";
  import ChevronRight from "@lucide/svelte/icons/chevron-right";
  import Building2 from "@lucide/svelte/icons/building-2";
  import type { Map as MaplibreMap } from "maplibre-gl";
  import { Button } from "#lib/components/ui/button/index.js";

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

  const INITIAL_COUNT = 5;
  let cardOpen = $state<boolean>(true);
  let showAll = $state<boolean>(false);
  const visibleBranches = $derived.by(() => {
    const sorted = [...branches].sort((a, b) =>
      a.type === "head" ? -1 : b.type === "head" ? 1 : 0,
    );
    return showAll ? sorted : sorted.slice(0, INITIAL_COUNT);
  });

  function flyToBranch(branch: Branch) {
    mapInstance?.flyTo({
      center: [branch.longitude, branch.latitude],
      zoom: 14,
      pitch: 50,
      bearing: 20,
      duration: 1500,
      essential: true,
    });
  }

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
    options={{ pitch: 50, bearing: 20 }}
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
              class="flex h-8 w-8 items-center justify-center rounded-full bg-red-500 text-white shadow-lg ring-2 ring-white dark:ring-zinc-900 hover:animate-bounce"
            >
              <Landmark class="h-3.5 w-3.5" />
            </div>
          {:else}
            <div
              class="flex h-9 w-9 items-center justify-center rounded-full bg-white text-white shadow-lg hover:animate-bounce"
            >
              <Avatar.Root>
                <Avatar.Image src={capitecLogo} alt="Capitec Bank" />
                <Avatar.Fallback>CB</Avatar.Fallback>
              </Avatar.Root>
            </div>
          {/if}
        </MarkerContent>

        <MarkerPopup class="overflow-hidden p-0 max-w-none">
          <div
            class={[
              "px-4 py-3",
              branch.type === "head" ? "bg-red-500" : "bg-sky-600",
            ].join(" ")}
          >
            <div class="flex items-center gap-3">
              <div
                class="flex h-9 w-9 shrink-0 items-center justify-center rounded-full {branch.type ===
                'head'
                  ? 'bg-red-500'
                  : 'bg-white'}"
              >
                {#if branch.type === "head"}
                  <Landmark class="h-4 w-4 text-white" />
                {:else}
                  <Avatar.Root class="h-8 w-8">
                    <Avatar.Image src={capitecLogo} alt="Capitec" />
                    <Avatar.Fallback class="bg-white text-xs text-white"
                      >CB</Avatar.Fallback
                    >
                  </Avatar.Root>
                {/if}
              </div>

              <div class="min-w-0 flex-1">
                <p
                  class="text-[10px] font-medium tracking-widest text-white/70"
                >
                  {branch.type === "head"
                    ? "Head Office"
                    : branch.type === "atm"
                      ? "ATM"
                      : "Branch"}
                </p>
                <h3
                  class="truncate text-sm font-semibold leading-tight text-white"
                >
                  {branch.name}
                </h3>
              </div>
            </div>
          </div>

          <div class="w-60 space-y-2.5 px-4 py-3">
            <div class="flex items-start gap-2.5">
              <span
                class={[
                  "flex h-5 w-5 shrink-0 items-center justify-center rounded",
                  branch.type === "head" ? " text-red-600" : " text-sky-600",
                ].join(" ")}
              >
                <MapPin class="h-3 w-3" />
              </span>
              <span class="text-xs leading-5 text-muted-foreground"
                >{branch.city}, {branch.province}</span
              >
            </div>

            {#if branch.phone}
              <div class="flex items-center gap-2.5">
                <span
                  class={[
                    "flex h-5 w-5 shrink-0 items-center justify-center rounded",
                    branch.type === "head" ? " text-red-600" : " text-sky-600",
                  ].join(" ")}
                >
                  <Phone class="h-3 w-3" />
                </span>
                <span class="text-xs text-muted-foreground">{branch.phone}</span
                >
              </div>
            {/if}

            <div class="flex items-start gap-2.5">
              <span
                class={[
                  "flex h-5 w-5 shrink-0 items-center justify-center rounded",
                  branch.type === "head" ? " text-red-600" : " text-sky-600",
                ].join(" ")}
              >
                <Clock class="h-3 w-3" />
              </span>
              <span class="text-xs leading-5 text-muted-foreground"
                >{branch.hours}</span
              >
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

  <div class="absolute top-4 left-4 z-10 w-72">
    <div
      class="overflow-hidden rounded-xl border bg-background/95 shadow-xl backdrop-blur-sm"
    >
      <Button
        type="button"
        onclick={() => (cardOpen = !cardOpen)}
        class="flex w-full items-center justify-between px-4 py-3 transition-colors hover:bg-muted/50"
      >
        <div class="flex items-center gap-2">
          <div class="flex h-6 w-6 items-center justify-center rounded-md">
            <Building2 class="h-3.5 w-3.5 text-sky-600" />
          </div>
          <span class="text-sm font-semibold">Branches</span>
        </div>
        <ChevronDown
          class="h-4 w-4 text-muted-foreground transition-transform duration-200 {cardOpen
            ? 'rotate-180'
            : ''}"
        />
      </Button>

      {#if cardOpen}
        <div class="border-t">
          <ul class="divide-y divide-border">
            {#each visibleBranches as branch (branch.id)}
              <li>
                <Button
                  type="button"
                  onclick={() => flyToBranch(branch)}
                  class="flex w-full items-center gap-3 px-8 py-8 text-left transition-colors hover:bg-muted/50 cursor-pointer"
                >
                  <div
                    class={[
                      "flex h-7 w-7 shrink-0 items-center justify-center rounded-full",
                      branch.type === "head"
                        ? "text-white bg-red-600"
                        : branch.type === "atm"
                          ? "text-amber-600 "
                          : "text-sky-600",
                    ].join(" ")}
                  >
                    {#if branch.type === "head"}
                      <Landmark class="h-3.5 w-3.5" />
                    {:else}
                      <Avatar.Root class="h-4 w-4">
                        <Avatar.Image src={capitecLogo} alt="Capitec" />
                        <Avatar.Fallback class="bg-white text-xs text-white"
                          >CB</Avatar.Fallback
                        >
                      </Avatar.Root>
                    {/if}
                  </div>

                  <div class="min-w-0 flex-1">
                    <p class="truncate text-xs font-medium">{branch.name}</p>
                    <p class="truncate text-xs text-muted-foreground">
                      {branch.city}, {branch.province}
                    </p>
                  </div>
                </Button>
              </li>
            {/each}
          </ul>

          {#if branches.length > INITIAL_COUNT}
            <div class="border-t px-4 py-2.5">
              <button
                type="button"
                onclick={() => (showAll = !showAll)}
                class="flex w-full items-center justify-center gap-1.5 py-1 text-xs font-medium text-sky-600 transition-colors hover:text-sky-700"
              >
                {showAll
                  ? "Show less"
                  : `Show ${branches.length - INITIAL_COUNT} more`}
                <ChevronDown
                  class="h-3.5 w-3.5 transition-transform duration-200 {showAll
                    ? 'rotate-180'
                    : ''}"
                />
              </button>
            </div>
          {/if}
        </div>
      {/if}
    </div>
  </div>
</div>
