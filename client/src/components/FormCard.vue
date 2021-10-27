<template>
    <div class="base" @click="showFields">
        <img width="120" height="200" alt="form image" :src="'data:image/png;base64,'.concat(form.form_img)"/>
        <br/>
        <span class="formName">{{ form.name }}</span>
    </div>

    <!-- should've been a separate component but eh.. -->
    <div id="fields" class="modalBody w3-modal">
        <div class="w3-modal-content">
            <div class="w3-container">
                    <span @click="hideFields"
                        class="w3-button w3-display-topright close">
                        &times;
                    </span>

                <Fields :form="form"/>

            </div>
        </div>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import Fields from "@/components/Fields.vue";
import Form from "@/models/Form";

export default defineComponent({
    name: "FormCard",
    components: {
        Fields
    },
    props: {
        form: Form
    },
    methods: {
        showFields() {
            document.getElementById('fields')!.style.display = 'block';
        },
        hideFields() {
            document.getElementById('fields')!.style.display = 'none';
        },
    }
});
</script>

<style scoped>
/* form */
.base {
    background-color: rgba(1, 1, 1, 0);
    cursor: pointer;

    padding: 10px 0 0 10px;
    margin: 10px;
    border: 1px rgba(1, 1, 1, 0) solid;
}

.base:hover {
    background-color: rgba(200, 200, 200, 50);
    border: 1px #34A853 solid;
}

.formName {
    color: #1c1c2e
}

/* fields */
@keyframes popup {
    0% {
        transform: scale(1);
    }
    50% {
        transform: scale(1.4);
    }
    60% {
        transform: scale(1.1);
    }
    70% {
        transform: scale(1.2);
    }
    80% {
        transform: scale(1);
    }
    90% {
        transform: scale(1.1);
    }
    100% {
        transform: scale(1);
    }
}

.modalBody {
    border-radius: 5px;

    animation-name: popup;
    animation-duration: 0.25s;
}

.close {
    color: #FFFFFF;
    background-color: #EA4335;
}
</style>