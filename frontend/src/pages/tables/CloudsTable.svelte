<script lang="ts">
//  import myData from '$site/data/data'
  export let data=[]
  import { DataHandler, check, Datatable, Th, ThFilter } from '@vincjo/datatables'
  import { onMount } from 'svelte'


 function setGeneralCheckbox(){
  let countSelectedTrue = data.filter(item => item.select === true).length;
  console.log(data.length,countSelectedTrue)
  if (countSelectedTrue == data.length){
  document.getElementById("my-checkbox").indeterminate = false
  document.getElementById("my-checkbox").checked = true
  }else if(countSelectedTrue == 0 ){
  document.getElementById("my-checkbox").indeterminate = false
  document.getElementById("my-checkbox").checked = false
  }else{
  document.getElementById("my-checkbox").indeterminate = true
  document.getElementById("my-checkbox").checked = false
  }
 } 



  const handler = new DataHandler(data, { rowsPerPage: 50 })
  handler.initializeSelection = () => {
  data.forEach((row) => {
    if (row.select) {
      console.log(row.select)
      handler.select(row.cloud);
  handler.selectAll({ selectBy: 'cloud' })
    }
  });
};
let aaa=false
  $: {
  data, handler.setRows(data)
  }
  const rows = handler.getRows()
  const selected = handler.getSelected()
  function getSS(fn){
    return $selected.includes(fn)
  }
  const isAllSelected = handler.isAllSelected()
  onMount(() => {
    handler.initializeSelection();
  handler.selectAll({ selectBy: 'cloud' })
  setGeneralCheckbox()
    //
  });
  let selectFilter="all";
  function filterBySelect(){
    if (selectFilter == "true" || selectFilter == "false"){
    handler.filter(selectFilter, 'select', check.isEqualTo)
    }else{
    handler.filter('', 'select', check.isEqualTo)
  }}
  filterBySelect()
</script>

<Datatable {handler}>
  <table>
      <thead>
          <tr>
              <Th {handler} class="selection" orderBy="select" >
                <input type="checkbox" class="" id="my-checkbox" 
                      on:click={(event) => {
                        //handler.selectAll({ selectBy: 'cloud' })
                        if (event.target.checked){
                          //console.log("fug")
                          data.forEach((element) => element.select=true)
                          setGeneralCheckbox()
                          //console.log(data)
                        }else{
                          data.forEach((element) => element.select=false)
                          setGeneralCheckbox()
                        }}}
                
                
                />
                 
              </Th>
              <Th {handler} orderBy="cloud">cloud</Th>
          </tr>
          <tr>
             <!-- <ThFilter {handler} filterBy="select" comparator={check.isEqualTo} class="selection" /> -->
             <th class="text-left">
              <select bind:value={selectFilter}
               on:change={(event) => {
                    if (event.target.value == "true" || event.target.value == "false"){
                    handler.filter(event.target.value, 'select', check.isEqualTo)
                    }else{
                    handler.filter('', 'select', check.isEqualTo)
                    }
                        //handler.selectAll({ selectBy: 'cloud' })
                        //selectFilter=false
                        
                        }}
              class="select select-xs">
                <option selected>all</option>
                <option>true</option>
                <option>false</option>
              </select>
            </th>
             <!-- <ThFilter {handler} filterBy="id" comparator={check.isEqualTo} /> -->
              <ThFilter {handler} filterBy="cloud" />
          </tr>
      </thead>
      <tbody>
          {#each $rows as row}
              <tr  class:active={$selected.includes(row.cloud)}>
                  <td  class="selection text-left">
                      <input
                          type="checkbox"
                          id={row.cloud}
                          on:click={() => {
                            handler.select(row.cloud)
                            let selected=data.find(objet => objet.cloud === row.cloud)
                            //document.getElementById(row.cloud).checked=row.select
                            selected.select=!selected.select
                            console.log(data)
                            filterBySelect()
                            setGeneralCheckbox()
                          }}
                          bind:checked={row.select}

                      />
                  </td>
                  <td>{@html row.cloud}</td>
              </tr>
          {/each}
      </tbody>
  </table>

                           <!-- checked={$selected.includes(row.cloud)} */ <tr class:active={$selected.includes(row.cloud) || row.select}>   bind:checked={row.select}  -->
</Datatable>

<style>
  thead {
      background: #fff;
  }
  thead th.selection {
      width: 48px;
      padding-left: 8px;
      border-bottom: 1px solid #e0e0e0;
  }
  tbody td {
      border: 1px solid #f5f5f5;
      padding: 4px 20px;
  }
  tbody tr {
      transition: all, 0.2s;
  }
  tbody tr:hover {
      background: #f5f5f5;
  }
  tbody tr.active {
      background: var(--primary-lighten-1);
  }
  tbody tr.active:hover {
      background: var(--primary-lighten-2);
  }
  td :global(b) {
      font-weight: normal;
      color: #bdbdbd;
      font-family: JetBrains;
      font-size: 11px;
  }
  td.selection {
      padding-left: 24px;
  }
</style>


