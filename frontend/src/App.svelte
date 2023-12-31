<script>

  import Sidebar from './Sidebar.svelte';
  import { iconDebug, iconSettings, iconProjects, softIconServers} from './icons.js';
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
  		{name: "Servers", icon: softIconServers , page: PageServers },
		{name: "Settings", icon: iconSettings , page: PageSettings },
		{name: "Debug", icon: iconDebug , page: PageDebug }
		];

</script>

<main>


ddd
<Sidebar mainButtons={mainButtons}/>



















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
