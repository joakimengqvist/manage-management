/* eslint-disable @typescript-eslint/no-explicit-any */
import { Button, Typography, Input } from "antd"
import { replaceUnderscore } from "../../helpers/stringFormatting";

const { Text, Title } = Typography;
const { TextArea } = Input;

type CreateNoteProps = {
    type: string
    title: string
    onTitleChange: (event : any) => void
    note: string
    onNoteChange: (event : any) => void
    onClearNoteFields: () => void
    onSubmit: () => void
}

const CreateNote = (props : CreateNoteProps) => {
    const { type, title, onTitleChange, note, onNoteChange, onClearNoteFields, onSubmit } = props;
    return (<>
        <Title level={5} style={{margin: '0px'}}>Create new {replaceUnderscore(type)} note</Title>
        <Text strong style={{lineHeight: 2.3}}>Title</Text>
        <Input value={title} onChange={onTitleChange} />
        <Text strong style={{lineHeight: 2.3}}>Note</Text>
        <TextArea value={note} onChange={onNoteChange}/>
        <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px', marginTop: '12px'}}>
          <Button onClick={onClearNoteFields} type="link">Clear</Button>
          <Button type="primary" disabled={!note || !title} onClick={onSubmit}>Submit</Button>
        </div>
        </>
    )
}

export default CreateNote;