<script>
  import { GetServers } from '../../wailsjs/go/main/App';
  let config="" 

  function getservers() {
    GetServers().then((result) => {
      config = result;
    });
  }

getservers()

</script>

{config}