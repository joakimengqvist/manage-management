/* eslint-disable @typescript-eslint/no-explicit-any */
import { Typography } from "antd";
import { Notes } from '../../interfaces/notes';
import { replaceUnderscoreAndCapitalize } from '../../helpers/stringFormatting';
import Note from "./Note";

const { Title } = Typography

interface NoteListProps {
    notes: Array<Notes>
    type: string
    userId: string
    generalized?: boolean
}

const NoteList = (props: NoteListProps) => {
    const {notes, type, userId, generalized} = props;
    return (<>
        <Title level={5}>{replaceUnderscoreAndCapitalize(type)} notes</Title>
        {notes.length > 0 && notes.map((note : any) => (
        <Note 
            note={note} 
            type={type} 
            userId={userId}
            generalized={generalized}
            />
        ))}
      </>)
}

export default NoteList