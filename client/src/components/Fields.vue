<template>
    <form @submit.prevent="makeForm">
        <table>
            <tr v-for="field in form1.fields" :key="field">
                <td>
                    <label>{{ field.name }} : </label>
                </td>
                <td>
                    <input v-if="isTextField(field)" type="text" v-model="field.content.text" required/>
                    <textarea v-if="isMultiLinedTextField(field)" v-model="field.content.text" required>
                </textarea>
                    <select id="s" v-if="isSelectionField(field)" v-model="field.content.selection">
                        <option v-for="selection in getKeys(field.content.selections)" :key="selection">
                            {{ selection }}
                        </option>
                    </select>
                </td>

            </tr>
        </table>
        <input type="submit" value="Make form"/>
    </form>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import Field, {FieldType} from "@/models/Field/Field";
import Form from "@/models/Form";

export default defineComponent({
    name: "Fields",
    props: {
        form: Form
    },
    data() {
        return {
            form1: this.getFormWithFieldsFromStorage(),
        }
    },
    mounted() {
        // this.getFieldsValues();  
    },
    methods: {
        isTextField(field: Field): boolean {
            return field.field_type === FieldType.TextField;
        },
        isSelectionField(field: Field): boolean {
            return field.field_type === FieldType.SelectionField;
        },
        isMultiLinedTextField(field: Field): boolean {
            return field.field_type === FieldType.MultiLinedTextField;
        },
        getKeys(content: { string: number }): string[] {
            const keys: string[] = [];

            for (const c in content) {
                keys.push(c);
            }
            return keys;
        },
        makeForm(): void {
            this.setFieldsValues();
            this.form.generateForm();
        },
        getFormWithFieldsFromStorage(): Form {
            const form1 = new Form(this.form.name, this.form.fields, this.form.form_img);

            for (let i in form1.fields) {
                const storageName = `${form1.name}_${form1.fields[i].name}`;
                const fieldValue = localStorage.getItem(storageName);
                if (fieldValue != null && form1.fields[i].savable) {
                    form1.fields[i].content.text = fieldValue;
                }
            }

            return form1;
        },
        setFieldsValues(): void {
            for (const field of this.form1.fields) {
                const storageName = `${this.form1.name}_${field.name}`;
                if (field.savable) {
                    localStorage.setItem(storageName, field.content.text as string);
                }
            }
        }
    }
});
</script>

<style scoped>

</style>
