import TextField from "./TextField";
import SelectionField from "./SelectionField";
import MultiLinedTextField from "./MultiLinedTextField";

export enum FieldType {
    TextField = 0,
    SelectionField,
    MultiLinedTextField
}

interface Field {
    name: string;
    content: any;
    field_type: FieldType;
    savable: boolean;
    setContent(content: any): void;
}

export function createField(name: string, content: any, fieldType: FieldType, savable: boolean): Field {
    switch (fieldType) {
        case FieldType.TextField:
            return new TextField(name, <string>content, savable);
        case FieldType.SelectionField:
            return new SelectionField(name, <{ string: number }>content, savable);
        case FieldType.MultiLinedTextField:
            return new MultiLinedTextField(name, <string>content, savable);
    }
}

export default Field;
