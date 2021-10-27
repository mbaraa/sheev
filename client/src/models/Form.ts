import Field, {createField} from "@/models/Field/Field";
import {backendAddress} from "@/main";

class Form {
    private static httpClient = new XMLHttpRequest();

    name: string;
    fields: Field[];
    form_img: string; // in base64

    constructor(name: string, fields: Field[], form_img: string) {
        this.name = name;
        this.fields = fields;
        this.form_img = form_img;
    }

    /**
     * generateForm sends the current working form's data to the backend,
     * and awaits for the generated image to be generated/sent back
     * */
    public async generateForm(): Promise<void> {
        await fetch(`${backendAddress}/forms/gen/`, {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                "name": this.name,
                "fields": this.fields
            })
        })
            .then(resp => resp.json())
            .then(data => {
                const a: HTMLAnchorElement = document.createElement("a");
                a.href = `data:image/png;base64,${data["img"]}`;
                a.download = `${this.name} ${new Date().toDateString()}`;
                a.click();
            })
            .catch(err => console.log(err))
    }

    public static getAllForms(): Form[] {
        Form.httpClient.open("GET", `${backendAddress}/forms/all/`, false)
        Form.httpClient.send();
        const forms = <Form[]>JSON.parse(Form.httpClient.responseText)["forms"];

        for (let i = 0; i < forms.length; i++) {
            forms[i] = new Form(forms[i].name, forms[i].fields, forms[i].form_img);
        }

        return forms;
    }

    /**
     * loadFormFromServer loads form's elements from the server
     * */
    public static loadFormFromServer(formName: string): Form {
        Form.httpClient.open("GET", `${backendAddress}/forms/single/${this.name}`, false);
        Form.httpClient.send();

        try {
            const resp = JSON.parse(Form.httpClient.responseText);

            const form: Form = new Form(
                formName,
                <Field[]>resp["fields"],
                <string>resp["form_img"],
            );

            for (let i = 0; i < resp["fields"].length; i++) {
                form.fields[i] = createField(resp["fields"][i].name, resp["fields"][i].content, resp["fields"][i].field_type, resp["fields"][i].savable);
            }

            return form;

        } catch (e) {
            return new Form("", [], "");
        }
    }


}

export default Form;
