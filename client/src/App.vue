<template>
    <div style="width: 100%; height: 100vh">
        <TopBar/>
        <Forms/>
    </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import TopBar from "./components/TopBar.vue"
import Forms from "@/components/Forms.vue";

export default defineComponent({
    name: 'App',
    components: {
        Forms,
        TopBar,
    },
});
</script>

<style>
#app {
    height: 100%;
    font-family: Avenir, sans-serif;
}
</style>
