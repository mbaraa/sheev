import Field, {FieldType} from "./Field";

class TextField implements Field {
    name: string;
    content: any;
    field_type: FieldType;
    savable: boolean;

    constructor(name: string, content: string, savable: boolean) {
        this.name = name;
        this.content = content;
        this.field_type = FieldType.TextField;
        this.savable = savable;
    }

    public setContent(content: any): void {
        this.content.text = <string>content;
    }
}

export default TextField;