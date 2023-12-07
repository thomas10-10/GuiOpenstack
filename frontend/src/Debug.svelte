         <script>
  import Counter from './lib/Counter.svelte'
  import { Greet } from '../wailsjs/go/main/App';

  let name = "";
  let greeting = "";

  function greet() {
      Readconf().then((result) => {
      greeting = result;
    });


  }

	 </script>
fucjjjjj


	 <li>
            <a class="flex cursor-pointer   items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-800 group"  on:click={() => toggleTabs(2)} >
	    <svg class="h-8 w-8 text-blue-500"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round">  <rect x="2" y="2" width="20" height="8" rx="2" ry="2" />  <rect x="2" y="14" width="20" height="8" rx="2" ry="2" />  <line x1="6" y1="6" x2="6.01" y2="6" />  <line x1="6" y1="18" x2="6.01" y2="18" /></svg>
                             <span class="flex-1 ms-3 whitespace-nowrap">Debug</span>
               <span class="inline-flex items-center justify-center px-2 ms-3 text-sm font-medium text-gray-800 bg-gray-100 rounded-full dark:bg-gray-700 dark:text-gray-300">Beta</span>
            </a>
         </li>

