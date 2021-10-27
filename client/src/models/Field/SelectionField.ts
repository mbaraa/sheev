import Field, {FieldType} from "./Field";

class SelectionField implements Field {
    name: string;
    content: any; // selectionName: order
    field_type: FieldType;
    savable: boolean;

    constructor(name: string, content: { string: number }, savable: boolean) {
        this.name = name;
        this.content = content;
        this.field_type = FieldType.SelectionField;
        this.savable = savable;
    }

    public setContent(content: any): void {
        this.content.selection = <string>content;
    }
}

export default SelectionField;
