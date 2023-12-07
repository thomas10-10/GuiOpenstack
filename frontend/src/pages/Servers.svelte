<script>
  import Table from '../Table.svelte';
  //import HTable from '../HandsonTable.svelte'

	let data = [
		['','Tesla','Volvo','Toyota','Ford'],
		['2019','Tesla','Volvo','Toyota','Ford'],
		['2020','Tesla','Volvo','Toyota','Ford'],
		['2021','Tesla','Volvo','Toyota','Ford']
	]
  import { GetServers } from '../../wailsjs/go/main/App';
  let config="" 
  let error="" 

  function getservers() {
    GetServers().then((result) => {
      if (result["Type"] === 'Message') {
        error = result;
      }else{
        config = result["Array"];
      } 
    });
  }

getservers()

</script>



<div class="p-4 sm:ml-64">
<div class="collapse bg-base-200">
  <input type="checkbox" class="peer" /> 
  <div class="collapse-title bg-primary text-primary-content peer-checked:bg-secondary peer-checked:text-secondary-content">
    Click me to show/hide content
  </div>
  <div class="collapse-content bg-primary text-primary-content peer-checked:bg-secondary peer-checked:text-secondary-content"> 
    <p>hello</p>
  </div>
</div>



{#if error !== '' }
<div role="alert" class="alert alert-error">
  <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
  <span>{error}</span>
</div>
{/if}

{config}

{#if config == '' }
  <progress class="progress w-56"></progress>
{/if}
<Table arrayTR={config}/>

</div>