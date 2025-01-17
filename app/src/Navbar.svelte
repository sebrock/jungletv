<script lang="ts">
    import { link } from "svelte-navigator";
    import { apiClient } from "./api_client";
    import { darkMode, rewardAddress, rewardReceived } from "./stores";
    import { fade, fly } from "svelte/transition";
    import { globalHistory } from "svelte-navigator";
    import Toggle from "svelte-toggle";
    const historyStore = { subscribe: globalHistory.listen };

    let navbarOpen = false;

    historyStore.subscribe(() => {
        navbarOpen = false;
    });

    function setNavbarOpen() {
        navbarOpen = !navbarOpen;
    }

    let rAddress = "";
    let lastReward = "";
    let hideRewardTimeout: number;

    rewardAddress.subscribe((address) => {
        rAddress = address;
    });

    rewardReceived.subscribe((reward) => {
        clearTimeout(hideRewardTimeout);
        lastReward = reward;
        hideRewardTimeout = setTimeout(() => (lastReward = ""), 7000);
    });

    async function copyAddress(address: string) {
        try {
            await navigator.clipboard.writeText(address);
        } catch (err) {
            console.error("Failed to copy!", err);
        }
    }
</script>

<nav
    class="top-0 fixed z-50 {navbarOpen
        ? 'h-auto'
        : 'h-16'} w-full flex flex-wrap items-center justify-between px-2 py-3 navbar-expand-lg bg-white shadow dark:bg-gray-950 dark:text-gray-300"
>
    <div class="container max-w-none w-full px-4 mx-auto flex flex-wrap items-center justify-between">
        <div class="w-full relative flex justify-between lg:w-auto lg:static lg:block lg:justify-start">
            <a
                use:link
                class="text-blueGray-700 text-sm font-bold leading-relaxed inline-block mr-4 whitespace-nowrap uppercase"
                href="/"
            >
                <img src="/assets/brand/logo.svg" alt="JungleTV" class="h-11 -mb-2" />
            </a>
            <button
                class="cursor-pointer text-xl leading-none px-3 py-1 border border-solid border-transparent rounded bg-transparent block lg:hidden outline-none focus:outline-none"
                type="button"
                on:click={setNavbarOpen}
            >
                <i class="fas fa-bars" />
            </button>
        </div>
        <div class="lg:flex flex-grow items-center {navbarOpen ? 'block mt-4' : 'hidden'}">
            <ul class="flex flex-col lg:flex-row list-none mr-auto">
                <li class="flex items-center">
                    {#if rAddress !== ""}
                        <span class="text-xs text-gray-500 mt-2 mb-4 lg:mt-0 lg:mb-0">
                            Rewarding <img
                                src="https://monkey.banano.cc/api/v1/monkey/{rAddress}?format=png"
                                alt="&nbsp;"
                                title="Click to copy: {rAddress}"
                                class="inline h-9 -mt-5 -mb-4 -ml-1 -mr-1 cursor-pointer"
                                on:click={() => copyAddress(rAddress)}
                            />
                            <span
                                class="font-mono cursor-pointer"
                                title="Click to copy: {rAddress}"
                                on:click={() => copyAddress(rAddress)}
                            >
                                {rAddress.substr(0, 16)}
                            </span>
                        </span>
                    {/if}
                    {#if lastReward !== ""}
                        <span
                            class="text-sm text-gray-700 bg-yellow-200 ml-5 p-1 rounded"
                            in:fly={{ x: 200, duration: 1000 }}
                            out:fade
                        >
                            Received <span class="font-bold">{apiClient.formatBANPrice(lastReward)} BAN</span>!
                        </span>
                    {/if}
                    <!-- <a
                        class="hover:text-blueGray-500 text-blueGray-700 px-3 py-2 flex items-center text-xs uppercase font-bold"
                        use:link
                        href="/rewards/address"
                    >
                        <i
                            class="text-blueGray-400 fas fa-coins text-lg leading-lg mr-2"
                        />
                        Earn rewards
                    </a> -->
                </li>
            </ul>
            <ul class="flex flex-col lg:flex-row list-none lg:ml-auto">
                <li class="flex items-center">
                    <div class="lg:mb-0 ml-3 mb-3 flex flex-row">
                        <i class="fas fa-sun text-lg leading-lg mr-2 text-gray-500" />
                        <Toggle
                            bind:toggled={$darkMode}
                            hideLabel
                            label="Toggle dark mode"
                            toggledColor="#6b7280"
                            untoggledColor="#6b7280"
                        />
                        <i class="fas fa-moon text-lg leading-lg ml-2 text-gray-500" />
                    </div>
                </li>
                <li class="flex items-center">
                    <a
                        class="dark:bg-gray-900 dark:text-gray-300 text-gray-700 text-xs font-bold uppercase px-4 py-2 rounded shadow hover:shadow-lg hover:bg-yellow-200 dark:hover:bg-yellow-900 outline-none focus:outline-none lg:mr-1 lg:mb-0 ml-3 mb-3 ease-linear transition-all duration-150"
                        use:link
                        href="/about"
                    >
                        <i class="fas fa-info text-lg leading-lg mr-2" />
                        What is this?
                    </a>
                </li>

                <li class="flex items-center">
                    <a
                        class="dark:bg-gray-900 dark:text-purple-500 text-purple-700 text-xs font-bold uppercase px-4 py-2 rounded shadow hover:shadow-lg hover:bg-yellow-200 dark:hover:bg-yellow-900 outline-none focus:outline-none lg:mr-1 lg:mb-0 ml-3 mb-3 ease-linear transition-all duration-150"
                        use:link
                        href="/rewards/address"
                    >
                        <i class="fas fa-coins text-lg leading-lg mr-2" />
                        Earn rewards
                    </a>
                </li>

                <li class="flex items-center">
                    <a
                        class="dark:bg-yellow-600 bg-yellow-400 text-white text-xs font-bold uppercase px-4 py-2 rounded shadow hover:shadow-lg hover:bg-yellow-500 dark:hover:bg-yellow-500 outline-none focus:outline-none lg:mr-1 lg:mb-0 ml-3 mb-3 ease-linear transition-all duration-150"
                        use:link
                        href="/enqueue"
                    >
                        <i class="fas fa-plus" /> Enqueue video
                    </a>
                </li>
            </ul>
        </div>
    </div>
</nav>
