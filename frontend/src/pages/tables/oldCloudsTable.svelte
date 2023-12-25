



  <RevoGrid  resize="true" {columns} theme="material"/>


let columnso=[
{
  sort: 0,
  width: "15%",
  id: "select",
  //name: h('input', {
  //      id: "selectAllCloud",
  //      type: "checkbox",
  //      checked: selectAllcloud ,
	 //  			onChange: () => {
  //          selectAllcloud=!selectAllcloud
  //          
	 // },
  //      }
  //    ),


 name: html('<a onclick="() => ;">v</a> <a>x</a> '),
 // formatter: (cell,row) => {
 //   let getCloud=cloudNamesDict.find(objet => objet.cloud === row.cells[1].data)
 //   return html(`<input type="checkbox" checked=${getCloud.select} onChange="() => console.log(cloudNamesDict);getCloud.select=!getCloud.select"  />`)
 // }   
  formatter: (cell, row, index) => {
    let getCloud=cloudNamesDict.find(objet => objet.cloud === row.cells[2].data)
    return h('input', {
          type: "checkbox",
          checked: getCloud.select,
	   			onChange: () => {
            aa=false
            getCloud.select=!getCloud.select
            //this.checked=cell
	      		console.log(cloudNamesDict)
            columnso=columnso
            //cloudNamesDict=cloudNamesDict
	  }
		},)
  }
}
,{

  sort: 1,
  name: "select", 

  formatter: (cell, row, index) => {
    return cell.toString()
  }
},

,{
  sort: 1,
  name: "cloud"}
]
 
        <Grid  data={cloudNamesDict} columns={columnso} search=true    className={className} pagination={pagination} fixedHeader=true    />