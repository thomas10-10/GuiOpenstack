<script>

  import { iconDebug, iconSettings, iconProjects, iconServers } from './icons.js';
  import PageProjects from './pages/Projects.svelte';
  import PageServers from './pages/Servers.svelte';
  import PageSettings from './pages/Settings.svelte';
  import PageDebug from './pages/Debug.svelte';
  import logo from './assets/svelte.png'
  import Debug from './Debug.svelte';

    let openTab = 0;

  function toggleTabs(tabNumber) {
    openTab = tabNumber;
  }


  let mainButtons = [{name: "Projects", icon: iconProjects, page: PageProjects },
  		{name: "Servers", icon: iconServers , page: PageServers },
		{name: "Settings", icon: iconSettings , page: PageSettings },
		{name: "Debug", icon: iconDebug , page: PageDebug }
		];

</script>

<main>




<aside text-left  id="logo-sidebar" class="fixed top-0 left-0 z-40 w-64 h-screen pt-20 transition-transform -translate-x-full bg-white border-r border-gray-200 sm:translate-x-0 dark:bg-gray-800 dark:border-gray-700" aria-label="Sidebar">
   <div class="h-full px-3 text-left pb-4 overflow-y-auto bg-white dark:bg-gray-800">
      <ul class="space-y-2 font-medium" id="default-tab" data-tabs-toggle="#default-tab-content" role="tablist"  >
	
	{#each mainButtons as item, index}
	 <li>
            <a   class="flex items-center {openTab === index
        ? 'bg-gray-100 dark:bg-gray-700 '
        : ''} p-2 text-gray-900  cursor-pointer rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-800 group" on:click={() => toggleTabs(index)} >
   {@html item.icon} 
               <span class="flex-1 ms-3">{item.name}</span>
<span class="inline-flex items-center justify-center px-2 ms-3 text-sm font-medium text-gray-800 bg-gray-100 rounded-full dark:bg-gray-700 dark:text-gray-300">Beta</span>
            </a>
         </li>

	{/each}
       
<li>
            <a href="#" class=" hover:animate-pulse flex items-center p-2 h-12 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group       bg-gradient-to-br     from-pink-400    via-purple-400      to-indigo-400    hover:shadow-lg hover:scale-105 transition duration-500    ">
	    <svg class="h-9 w-9 text-blue-500"  width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">  <path stroke="none" d="M0 0h24v24H0z"/>  <path d="M6 5h12l3 5l-8.5 9.5a.7 .7 0 0 1 -1 0l-8.5 -9.5l3 -5" />  <path d="M10 12l-2 -2.2l.6 -1" /></svg>
                       <span class="ms-3 text-white font-bold">Donate</span>
            </a>
         </li>
	 <li>

	 </li>


      </ul>
   </div>
</aside>



 {#each mainButtons as item, index}
  {#if openTab === index }
  <svelte:component this={item.page} />

  {/if}
 {/each}















</main>

<style>
  :root {
    font-family: 'Nunito', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
      background-color: rgba(27, 38, 54, 1);
    text-align: center;
    color: white;
  }

  main {
    text-align: center;
    padding: 1em;
    margin: 0 auto;
    height: 100vh;
  }

  img {
    height: 16rem;
    width: 16rem;
  }

  h1 {
    @apply text-3xl font-bold underline text-emerald-700;
    /*color: #ff3e00;
    text-transform: uppercase;
    font-size: 4rem;
    font-weight: 100;*/
    line-height: 1.1;
    margin: 2rem auto;
    max-width: 14rem;
  }

  p {
    max-width: 14rem;
    margin: 1rem auto;
    line-height: 1.35;
  }

  
@font-face {
    font-family: "Nunito";
    font-style: normal;
    font-weight: 400;
    src: local(""),
    url("assets/fonts/nunito-v16-latin-regular.woff2") format("woff2");
}

  @media (min-width: 480px) {
    h1 {
      max-width: none;
    }

    p {
      max-width: none;
    }
  }


</style>
