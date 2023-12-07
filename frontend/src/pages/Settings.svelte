<script>
  import { Readconf, Writeconf } from '../../wailsjs/go/main/App';
  let config="" 

  function readconf() {
    Readconf().then((result) => {
      config = result;
    });
  }
  function writeconf() {
    Writeconf(config);
  }
readconf()

</script>

<div class="p-4 sm:ml-64 h-full ">
 <!-- <hr> -->
<div class="mockup-code pb-5 pr-1 m-0 h-full " >
<textarea placeholder="empty" style=" resize: none; "  on:change={writeconf} bind:value={config}   class=" m-0 .p-0 w-full h-full  textarea  textarea-lg  bg-transparent  leading-normal" ></textarea>

</div>
</div>
