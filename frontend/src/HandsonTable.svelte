
<script>
	import jQuery from 'jquery';
    import { onMount } from 'svelte';

	import Handsontable from 'handsontable';
    import 'handsontable/dist/handsontable.full.min.css';
    export let data;

    let gridElement;

    let gridStatus = {
        isScriptLoaded: false,
        isStyleLoaded: false,
        isMounted: false,
        isInited: false
    }
    
    onMount(() => {
        gridStatus.isMounted = true;
        if (gridStatus.isScriptLoaded && gridStatus.isStyleLoaded) gridInit()
    })

    function scriptLoaded() {
        gridStatus.isScriptLoaded = true;
        if (gridStatus.isMounted && gridStatus.isStyleLoaded) gridInit()
    }

    function styleLoaded() {
        gridStatus.isStyleLoaded = true;
        if (gridStatus.isScriptLoaded && gridStatus.isMounted) gridInit()
    }

    function gridInit() {
        if (!gridStatus.isInited) {
            gridStatus.isInited = true;
            new Handsontable(gridElement,{
                data:data,
                rowHeaders:true,
                colHeaders:true
            });
        } 
    }
    gridInit()
</script>
{@html gridElement}
<div bind:this={gridElement}></div>