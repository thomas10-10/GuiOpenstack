<script>

  import { iconLog, iconOkey, iconCloud, iconPlay } from '../icons.js';
  import {EventsOn, EventsOff } from "../../wailsjs/runtime/runtime";
  import { onMount } from 'svelte';
  import Table from '../Table.svelte';
  import CloudsTable from './tables/CloudsTable.svelte';
  //import HTable from '../HandsonTable.svelte'
  import alasql from 'alasql';
  import { compileExpression } from 'filtrex';
  import { h, html } from "gridjs"

import { RowSelection }  from "gridjs/plugins/selection"
import "gridjs/dist/theme/mermaid.css";

let selectAllcloud=true;
import {
  filter,
  highlight,
  parse,
  test,
} from 'liqe';


	let data = [
		['','Tesla','Volvo','Toyota','Ford'],
		['2019','Tesla','Volvo','Toyota','Ford'],
		['2020','Tesla','Volvo','Toyota','Ford'],
		['2021','Tesla','Volvo','Toyota','Ford']
	]
  import { GetServers, GetCloudNames, GreetAsyncViaEvent, AsyncGetServersViaEvent } from '../../wailsjs/go/main/App';


  let config=[] 
  EventsOn("rcv:getServers", (msg) => { 
    //console.log(msg[0]["ID"])
    //config = config.concat(msg);
    //console.log(msg)
    msg.forEach(function(element,index) {
      let a={ name: element["Name"], id: element["Id"], cloud: element["Cloud"]}
      config = config.concat(a);
    });
    let configa= [
    { name: msg[0]["ID"], email: "john@example.com" },
    { name: "Mark", email: "mark@gmail.com" },
  ]

  })
   EventsOn("rcv:displayError", (msg) => { 
    errors = errors.concat(msg);
    error=msg[0]+msg[1]
  })
   EventsOn("rcv:updateStateRequestGetServers", (msg) => { 
    reqStates[msg[0]]=msg[1]
  })


  
  EventsOn("rcv:greet", (msg) => document.getElementById("result").innerText = msg)
  function greetAsyncViaEvent(){
  // noinspection JSIgnoredPromiseFromCall
    GreetAsyncViaEvent();
  }

//greetAsyncViaEvent()
//greetAsyncViaEvent()
//AsyncGetServersViaEvent("sogate-noe")
//AsyncGetServersViaEvent("sogate-noe2")
//AsyncGetServersViaEvent("stargate-int")
//AsyncGetServersViaEvent("sogate-int")







  function getServerList(){
    GetCloudNames().then((result) => {
      if (result["Level"] != 'ERROR') {
        for (let element of result["Array"]) {
             reqStates[element] = "begin"
              GetServers(element).then((result) => {
                if (result["Type"] === 'Message') {
                    error = result;
                }else{
                  //config = result["Array"];
                  config = config.concat(result["Array"]);
                  //config.push(...result["Array"]);
                  }
              });
            
        } 
        console.log("okeety");
      }
    })
  }






  function getCloudNames() {
    let response={}
    response["error"]=""
    response["cloudNames"]=[]
    GetCloudNames().then((result) => {
      if (result["Level"] === 'ERROR') {
        response['error'] = result;
      }else{
        response["cloudNames"]=result["Array"];
        console.log("get list clouds "+response["cloudNames"])
        cloudNames=[]
        cloudNames=response["cloudNames"]
        cloudNamesDict=[]
        cloudNames.forEach((element) =>{
          console.log("gggggggggggggg",element)
          cloudNamesDict = cloudNamesDict.concat({"select":true ,"cloud":element});
          //cloudNamesDict.push({"cloud": element})
        });
        
      } 
    });
    return response

  }



  let errors=[] 
  let error="" 
  let cloudNames=[]
  let cloudNamesDict=[]
  
  
  cloudNames=getCloudNames()["cloudNames"]

  let reqStates={}






function processClouds() {
    console.log("fff")
  const promises = cloudNames.map((cloud) => {
    console.log(cloud)
    return getservers(cloud).then((result) => {
      reqStates[cloud] = "okey";
      console.log(result);
      console.log("heyyyyyyyyyyyyy");
      // Faire quelque chose dès qu'une opération est terminée
    });
  });

  Promise.all(promises)
    .then(() => {
      console.log("Toutes les opérations sont terminées");
      // Effectuez d'autres actions après l'exécution de toutes les promesses
    })
    .catch((error) => {
      console.error("Au moins une opération a échoué :", error);
    });
}

async function executeReqs(){
    cloudNames.forEach((element) =>{  
      reqStates[element] = "begin"
      getservers(element)
  
    });
}



  function getservers(cloud) {
    GetServers(cloud).then((result) => {
      if (result["Type"] === 'Message') {
        error = result;
      }else{
        config = result["Array"];
      }
    });
  }

cloudNames=getCloudNames()["cloudNames"] 
//cloudNames.forEach((element) =>  reqStates[element] = "begin");
//executeReqs()
//getservers("stargate-int")
//processClouds()

import { RevoGrid } from "@revolist/svelte-datagrid";
import {defineCustomElements} from '@revolist/revogrid/loader';

let source = [{
      prop: "name",
      name: "First",
    },
    {
      prop: "details",
      name: "Second",
}];
let columns = [{
    name: "1ffff",
    details: "Item 1",
}];
let i=0

let data2 = [
    { name: "John", email: "john@example.com" },
    { name: "Mark", email: "mark@gmail.com" },
  ]
const className= {
    table: 'stripe hover table table-xs'
  }


function selectAllClouds(){
  cloudNamesDict.forEach((element) => element.select=true    );
  selectAllcloud=true
  //cloudNamesDict=cloudNamesDict
}
function deselectAllClouds(){
  cloudNamesDict.forEach((element) => element.select=false    );
  //cloudNamesDict=cloudNamesDict
  console.log(cloudNamesDict)
}

let aa=true;

     

function log(...args){
  console.log(...args)
}


let pagination= {
    limit: 20,
    summary: true
  }


data2 = [
    { name: "Johnyyyy", email: "john@example.com" },
    { name: "Joyka", email: "john@example.com" },
    { name: "Mark", email: "mark@gmail.com" },
  ]

let filters=""
let view=[]
function filtrexit(){
  view=[]
  const f = compileExpression(filters)
  config.forEach(function(element,index) {
    if ( f(element) ){
      view = view.concat(element);
    }
  });
}

let errorAlasql=""
function filterit(){
if (filters == ""){
  filters="select * from ?"
}
//select * from ? where name REGEXP "^t.*"
try {
view=alasql(filters,[config])
}catch (error) {

    console.error("Une erreur s'est produite lors de l'exécution de la requête :", error.message);
}





//view= [...filter(parse(filters), config)]
//config.forEach(function(element,index) {
//      if (filters == ""){
//        filters= "1 == 1"
//      }
//      if ( eval(filters) ){ 
//        let a={ name: element["name"], id: element["id"], cloud: element["cloud"]}
//        console.log("fuckkk")
//      view = view.concat(a);
//      }
//    });
  }




</script>


<nav class="bg-gray-800">
  <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
    <div class="relative flex h-16 items-center justify-between">
      <div class="absolute inset-y-0 left-0 flex items-center sm:hidden">
        <!-- Mobile menu button-->
        <button type="button" class="relative inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white" aria-controls="mobile-menu" aria-expanded="false">
          <span class="absolute -inset-0.5"></span>
          <span class="sr-only">Open main menu</span>
          <!--
            Icon when menu is closed.

            Menu open: "hidden", Menu closed: "block"
          -->
          <svg class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
          </svg>
          <!--
            Icon when menu is open.

            Menu open: "block", Menu closed: "hidden"
          -->
          <svg class="hidden h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <div class="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
        <div class="flex flex-shrink-0 items-center">
          <img class="h-8 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500" alt="Your Company">
        </div>
        <div class="hidden sm:ml-6 sm:block">
          <div class="flex space-x-4">
            <!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
            <a href="#" class="bg-gray-900 text-white rounded-md px-3 py-2 text-sm font-medium" aria-current="page">Dashboard</a>
            <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium">Team</a>
            <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium">Projects</a>
            <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium">Calendar</a>
          </div>
        </div>
      </div>
      <div class="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
        <button type="button" class="relative rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
          <span class="absolute -inset-1.5"></span>
          <span class="sr-only">View notifications</span>
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0" />
          </svg>
        </button>

        <!-- Profile dropdown -->
        <div class="relative ml-3">
          <div>
            <button type="button" class="relative flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800" id="user-menu-button" aria-expanded="false" aria-haspopup="true">
              <span class="absolute -inset-1.5"></span>
              <span class="sr-only">Open user menu</span>
              <img class="h-8 w-8 rounded-full" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="">
            </button>
          </div>

          <!--
            Dropdown menu, show/hide based on menu state.

            Entering: "transition ease-out duration-100"
              From: "transform opacity-0 scale-95"
              To: "transform opacity-100 scale-100"
            Leaving: "transition ease-in duration-75"
              From: "transform opacity-100 scale-100"
              To: "transform opacity-0 scale-95"
          -->
          <div class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none" role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabindex="-1">
            <!-- Active: "bg-gray-100", Not Active: "" -->
            <a href="#" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1" id="user-menu-item-0">Your Profile</a>
            <a href="#" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1" id="user-menu-item-1">Settings</a>
            <a href="#" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1" id="user-menu-item-2">Sign out</a>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Mobile menu, show/hide based on menu state. -->
  <div class="sm:hidden" id="mobile-menu">
    <div class="space-y-1 px-2 pb-3 pt-2">
      <!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
      <a href="#" class="bg-gray-900 text-white block rounded-md px-3 py-2 text-base font-medium" aria-current="page">Dashboard</a>
      <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium">Team</a>
      <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium">Projects</a>
      <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium">Calendar</a>
    </div>
  </div>
</nav>

<input type="checkbox" value="dark" class="toggle theme-controller"/>


<div class="p-4 sm:ml-64">



  
  <div class="join join-vertical lg:join-horizontal">
    <button class="btn join-item"  onclick="my_modal_4.showModal()" ><span>{@html iconCloud}</span>Select Clouds</button>
    <dialog id="my_modal_4" class="modal">
      <div class="modal-box">
        <form method="dialog">
          <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
        </form>

  <CloudsTable data={cloudNamesDict} />
      </div>
    </dialog>






  <div class="dropdown dropdown-hover join-item">
    <div tabindex="0" role="button" class="btn join-item">
        {@html iconPlay} GET<span class=" badge badge-primary">last: 2019-12-15:00:00:00</span> 
    </div>
    <div tabindex="0" class="dropdown-content z-[1]  shadow bg-base-100 rounded-box w-52">
      <div>sssssssssss</div>
      <div>sssssssssss</div>
    </div>
  </div>


  </div>




  <div   class=" join flex  p-2  cursor-pointer rounded-lg  group"  >
 <button class="btn join-item  btn-sm w-1  ">
        <svg class=" flex-shrink-0 h-4 w-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
    </button>



  <div class="dropdown dropdown-hover join-item">
    <div tabindex="0" role="button" class="">
      <div class="indicator">
        <span class="indicator-item badge badge-error">error</span> 
        <select class="select select-bordered join-item select-sm border-error">
          <option class="" selected>SQL</option>
          <!--  <option disabled>jsonpath</option> -->
        </select>
      </div>    
    </div>
    <div tabindex="0" class="dropdown-content z-[1]  shadow bg-base-100 rounded-box w-52">
      <div>sssssssssss</div>
    </div>
  </div>






    
    
    <textarea class="input input-bordered join-item input-sm w-full" placeholder="Select * from ?"/>

<button class="btn join-item btn-sm"></button>
  </div>



  <div class="join">
      <button class="join-item btn btn-error btn-sm w-1  ">
        <svg class=" flex-shrink-0 w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"> <path stroke="none" d="M0 0h24v24H0z"/> <path d="M9 9v-1a3 3 0 0 1 6 0v1" /> <path d="M8 9h8a6 6 0 0 1 1 3v3a5 5 0 0 1 -10 0v-3a6 6 0 0 1 1 -3" /> <line x1="3" y1="13" x2="7" y2="13" /> <line x1="17" y1="13" x2="21" y2="13" /> <line x1="12" y1="20" x2="12" y2="14" /> <line x1="4" y1="19" x2="7.35" y2="17" /> <line x1="20" y1="19" x2="16.65" y2="17" /> <line x1="4" y1="7" x2="7.75" y2="9.4" /> <line x1="20" y1="7" x2="16.25" y2="9.4" /></svg>
      </button>


    <select class="select select-bordered join-item select-sm">
      <option selected>SQL</option>
     <!--  <option disabled>jsonpath</option> -->
    </select>
        <textarea class="input input-bordered join-item input-sm" placeholder="Select * from ?"/>
     <!--   <input class="input input-bordered join-item input-sm" placeholder="WHERE ..."/> -->
     <!--  <span class="indicator-item badge badge-secondary">BETA</span> -->
    <button class="btn join-item btn-primary btn-sm w-1  ">
        <svg class=" flex-shrink-0 h-4 w-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
    </button>

     </div>


  <div>
    <label for="hs-leading-button-add-on-with-leading-and-leading" class="sr-only">Label</label>
    <div class="flex rounded-lg shadow-sm">
      <button type="button" class="p-1 flex-shrink-0 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-s-md border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none dark:focus:outline-none dark:focus:ring-1 dark:focus:ring-gray-600">
        <svg class=" flex-shrink-0 h-4 w-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
      <span class="">SQL</span></button>
      <input type="text" id="hs-leading-button-add-on-with-leading-and-leading" name="hs-leading-button-add-on-with-leading-and-leading" class="py-3 px-4 block w-full border-gray-200 shadow-sm rounded-0 text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none dark:bg-slate-900 dark:border-gray-700 dark:text-gray-400 dark:focus:ring-gray-600">
      <span class="px-4 inline-flex items-center min-w-fit rounded-e-md border border-s-0 border-gray-200 bg-gray-50 text-sm dark:bg-gray-700 dark:border-gray-700">
        <span class="text-sm text-gray-500 dark:text-gray-400">http://</span>
      </span>
    </div>
  </div>





  <div   class="flex items-center bg-gray-100 dark:bg-gray-700  p-2 text-gray-900  cursor-pointer rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-800 group"  >
  <input class="w-full   " />         <span class="flex-1 ms-3"></span>
<button class="btn btn-error btn-xs">Beta</button>
<button class="btn btn-error btn-xs">Beta</button>
  </div>

  <input type="text" placeholder="Type here" on:keyup={filterit} bind:value={filters} class="flex-5 	 input input-bordered input-primary " />
  <button class="flex-5	 btn btn-error">Error</button>

  <input type="text" placeholder="Type here" class="input input-bordered input-xs w-full max-w-xs" />
  <div id="result"></div>
  <button class="btn" onclick="my_modal_3.showModal()">{@html iconLog} </button>
  <dialog id="my_modal_3" class="modal">
    <div class="modal-box">
      <form method="dialog">
        <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
      </form>

      {#each Object.keys(reqStates) as key }
        <div class="align-middle"><span>{key}</span>
          {#if reqStates[key] == "begin" || reqStates[key] == "pending"   }
          <span class="loading loading-spinner text-primary loading-xs">load</span>
          {/if}
          {#if reqStates[key] == "okey"  }
          <span class="">✓</span>
          {/if}
          {#if reqStates[key] == "error"  }
          <span class="">X</span>
          {/if}


        </div>
      {/each}
      <p class="py-4"></p>
    </div>
  </dialog>

<div class="collapse  bg-base-200">
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


{#if config == '' }
  <progress class="progress w-56"></progress>
{/if}

</div>