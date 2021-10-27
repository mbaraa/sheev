<template>
    <header class="nav" :style="{
        backgroundColor: '#34A853',
    }">
        <a href="#">Sheev</a>

    </header>
</template>

<script>
import {defineComponent} from "vue";

export default defineComponent({
    name: "TopBar",
});
</script>

<style scoped>
.nav {
    font-size: 2em;
    font-weight: bold;

    position: sticky;
    padding: 10px;
}

a {
    text-decoration: none;
    color: #FFFFFF;
    display: inline;
}
</style>
