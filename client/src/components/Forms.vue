<template>
    <div v-for="form in forms" :key="form" class="forms">
        <FormCard :form="form"/>
    </div>
</template>

<script lang="ts">
import FormCard from "./FormCard.vue";
import Form from "@/models/Form";
import {defineComponent} from "vue";

export default defineComponent({
    name: "Forms",
    components: {
        FormCard,
    },
    data() {
        return {
            forms: Form.getAllForms()
        }
    },
});
</script>

<style scoped>
.forms {
    display: inline-grid;
}
</style>