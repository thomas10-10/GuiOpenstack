<script lang="ts">
//  import myData from '$site/data/data'
  export let data=[]
  import { DataHandler, check, Datatable, Th, ThFilter } from '@vincjo/datatables'
  import { onMount } from 'svelte'

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
    //
  });
</script>

<Datatable {handler}>
  <table>
      <thead>
          <tr>
              <Th {handler} class="selection" orderBy="select" >
                  <input
                      type="checkbox"
                      on:click={(event) => {
                        handler.selectAll({ selectBy: 'cloud' })
                        if (event.target.checked){
                          console.log("fug")
                          data.forEach((element) => element.select=true)
                          console.log(data)
                        }else{

                          data.forEach((element) => element.select=false)
                        } 
                      }}
                      checked={$isAllSelected}
                      
                  />
              </Th>
              <Th {handler} orderBy="cloud">cloud</Th>
          </tr>
          <tr>
              <ThFilter {handler} filterBy="select" class="selection" />
             <!-- <ThFilter {handler} filterBy="id" comparator={check.isEqualTo} /> -->
              <ThFilter {handler} filterBy="cloud" />
          </tr>
      </thead>
      <tbody>
          {#each $rows as row}
              <tr class:active={$selected.includes(row.cloud)}>
                  <td class="selection">
                      <input
                          type="checkbox"
                          id={row.cloud}
                          on:click={() => {

                            handler.select(row.cloud)
                            let selected=data.find(objet => objet.cloud === row.cloud)
                            //document.getElementById(row.cloud).checked=row.select
                            selected.select=!selected.select
                            console.log(data)
                          }}
                          checked={$selected.includes(row.cloud)}

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


