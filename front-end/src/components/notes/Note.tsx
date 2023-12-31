/* eslint-disable @typescript-eslint/no-explicit-any */
import { Button, Divider, Input, notification, Popconfirm, Typography } from "antd";
import { Notes } from '../../interfaces/notes';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { NOTE_TYPE } from '../../enums/notes';
import { deleteIncomeNote } from "../../api/notes/income/delete";
import { deleteExpenseNote } from "../../api/notes/expense/delete";
import { deleteProjectNote } from "../../api/notes/project/delete";
import { DeleteOutlined, EditOutlined, SendOutlined, CloseCircleOutlined } from '@ant-design/icons';
import { deleteExternalCompanyNote } from "../../api/notes/externalCompany/delete";
import { formatDateTimeToYYYYMMDDHHMM } from "../../helpers/stringDateFormatting";
import { useEffect, useState } from "react";
import { updateExpenseNote } from "../../api/notes/expense/update";
import { updateIncomeNote } from "../../api/notes/income/update";
import { updateProjectNote } from "../../api/notes/project/update";
import { updateExternalCompanyNote } from "../../api/notes/externalCompany/update";
import { replaceUnderscore } from "../../helpers/stringFormatting";
import { PRIVILEGES } from "../../enums/privileges";
import { hasPrivilege } from "../../helpers/hasPrivileges";
import { deleteSubProjectNote } from "../../api/notes/subProject/delete";
import { updateSubProjectNote } from "../../api/notes/subProject/update";
import { useGetLoggedInUser } from "../../hooks";
import { updateProductNote } from "../../api/notes/product/update";
import { updateInvoiceNote } from "../../api/notes/invoice/update";
import { updateInvoiceItemNote } from "../../api/notes/invoiceItem/update";
import { deleteProductNote } from "../../api/notes/product/delete";
import { deleteInvoiceNote } from "../../api/notes/invoice/delete";
import { deleteInvoiceItemNote } from "../../api/notes/invoiceItem/delete";

const { Text, Link } = Typography;
const { TextArea } = Input;

interface NoteProps {
  type: string
  userId: string
  note: Notes
  generalized?: boolean
}

const Note = (props: NoteProps) => {
  const { type, generalized, userId, note } =  props;
  const [api, contextHolder] = notification.useNotification();
  const loggedInUser = useGetLoggedInUser();

  const [noteTitle, setNoteTitle] = useState('');
  const [noteBody, setNoteBody] = useState('');
  const [editing, setEditing] = useState(false);

  useEffect(() => {
    setNoteBody(note.note);
    setNoteTitle(note.title)
  }, [note])

  const onNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
  const onNoteBodyChange = (event : any) => setNoteBody(event.target.value);


    const noteFailedNotification = (message : string, action : string) => {
        api.error({
            message: `${action} project ${replaceUnderscore(type)} note failed`,
            description: message,
            placement: 'bottom',
            duration: 1.4
        });
    }

    const noteSuccessNotification = (response : any, action : string) => {
        if (response?.error) {
            noteFailedNotification(response.message, action)
        }
        api.info({
            message: response.message,
            placement: 'bottom',
            duration: 1.4
        });
    }

    const onClickDeleteNote = () => {
        if (type === NOTE_TYPE.expense) {
            deleteExpenseNote(userId, note.id)
            .then(response => {
                noteSuccessNotification(response, 'Deleted');
              })
              .catch((error) => {
                noteFailedNotification(error.toString(), 'Deleted')
              });
        }
        if (type === NOTE_TYPE.income) {
            deleteIncomeNote(userId, note.id)
            .then(response => {
                noteSuccessNotification(response, 'Deleted');
              })
              .catch((error) => {
                noteFailedNotification(error.toString(), 'Deleted')
              });
        }
        if (type === NOTE_TYPE.project) {
            deleteProjectNote(userId, note.id, note.author_id, note.project_id)
            .then(response => {
                noteSuccessNotification(response, 'Deleted');
              })
              .catch((error) => {
                noteFailedNotification(error.toString(), 'Deleted')
              });
        }
        if (type === NOTE_TYPE.sub_project) {
          deleteSubProjectNote(userId, note.id, note.author_id, note.sub_project_id)
          .then(response => {
              noteSuccessNotification(response, 'Deleted');
            })
            .catch((error) => {
              noteFailedNotification(error.toString(), 'Deleted')
            });
      }
        if (type === NOTE_TYPE.external_company) {
            deleteExternalCompanyNote(userId, note.id)
            .then(response => {
                noteSuccessNotification(response, 'Deleted');
              })
              .catch((error) => {
                noteFailedNotification(error.toString(), 'Deleted')
              });
        }
        if (type === NOTE_TYPE.product) {
          deleteProductNote(userId, note.id)
          .then(response => {
              noteSuccessNotification(response, 'Deleted');
            })
            .catch((error) => {
              noteFailedNotification(error.toString(), 'Deleted')
            });
      }
      if (type === NOTE_TYPE.invoice) {
        deleteInvoiceNote(userId, note.id)
        .then(response => {
            noteSuccessNotification(response, 'Deleted');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Deleted')
          });
    }
    if (type === NOTE_TYPE.invoice_item) {
      deleteInvoiceItemNote(userId, note.id)
      .then(response => {
          noteSuccessNotification(response, 'Deleted');
        })
        .catch((error) => {
          noteFailedNotification(error.toString(), 'Deleted')
        });
  }
    }

    const onSaveUpdateNote = () => {
      const noteAuthor = {
        id: loggedInUser.id,
        name: `${loggedInUser.firstName} ${loggedInUser.lastName}`,
        email: loggedInUser.email
      }
      if (type === NOTE_TYPE.expense) {
        if (!note.expense_id) return;
        updateExpenseNote(note.id, noteAuthor, note.expense_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }
    if (type === NOTE_TYPE.income) {
      if (!note.income_id) return;
        updateIncomeNote(note.id, noteAuthor, note.income_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }
    if (type === NOTE_TYPE.project) {
      if (!note.project_id) return;
        updateProjectNote(note.id, noteAuthor, note.project_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }
    if (type === NOTE_TYPE.sub_project) {
      if (!note.project_id) return;
        updateSubProjectNote(note.id, noteAuthor, note.project_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }
    if (type === NOTE_TYPE.external_company) {
      if (!note.external_company_id) return;
        updateExternalCompanyNote(note.id, noteAuthor, note.external_company_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }  
    if (type === NOTE_TYPE.product) {
      if (!note.product_id) return;
        updateProductNote(note.id, noteAuthor, note.product_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }   
    if (type === NOTE_TYPE.invoice) {
      if (!note.invoice_id) return;
        updateInvoiceNote(note.id, noteAuthor, note.invoice_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }   
    if (type === NOTE_TYPE.invoice_item) {
      if (!note.invoice_item_id) return;
        updateInvoiceItemNote(note.id, noteAuthor, note.invoice_item_id, noteTitle, noteBody)
        .then(response => {
            noteSuccessNotification(response.data, 'Updated');
          })
          .catch((error) => {
            noteFailedNotification(error.toString(), 'Updated')
          });
    }   
  }

    return (
      <div style={{width: '100%', border: '1px solid rgba(5, 5, 5, 0.06)', marginBottom: '8px', marginTop: '4px', borderRadius: '4px'}}>
        {contextHolder}
        <div style={{display: 'flex', justifyContent: 'space-between', alignItems: 'center', paddingTop: editing ? '0px' : '6px', margin: '0px 8px'}}>
        <div style={{width: '80%'}}>
          {editing ? (
            <Input style={{marginTop: '8px', marginBottom: '4px'}} value={noteTitle} onChange={onNoteTitleChange}/>
          ) : (
            <Text strong style={{margin: '0px', lineHeight: '32px'}}>{noteTitle}</Text>
          )}
        </div>
          <div>
          <Button onClick={() => setEditing(!editing)} type="link" style={{padding: '0px 8px 0px 0px'}}>
            {editing ? <CloseCircleOutlined /> : <EditOutlined />}
          </Button>
          {editing && (
            <Button type="link" onClick={onSaveUpdateNote}style={{padding: '0px 4px 0px 0px'}}><SendOutlined /></Button>
          )}
          {!editing && hasPrivilege(loggedInUser.privileges, PRIVILEGES.note_sudo) && (
            <Popconfirm
                placement="top"
                title="Are you sure?"
                description={`Do you want to delete note ${note.title}`}
                onConfirm={onClickDeleteNote}
                icon={<QuestionCircleOutlined twoToneColor="red" />}
                okText="Yes"
                cancelText="No"
            >
                <Button danger type="link" style={{padding: '0px 4px 0px 0px'}}><DeleteOutlined /></Button>
            </Popconfirm>
            )}
          </div>
        </div>
        <Divider style={{marginTop: '4px', marginBottom: '8px'}} />
        <div style={{margin: '0px 8px 8px 8px', minHeight: '50px'}}>
        {editing ? (
         <TextArea style={{marginTop: '8px', marginBottom: '4px'}} value={noteBody} onChange={onNoteBodyChange} />
        ): (
          <Text>{noteBody}</Text>
        )}
        </div>
        <Divider style={{ marginTop: '0px', marginBottom: '0px' }} />
        <div style={{display: 'flex', justifyContent: 'flex-end', flexDirection: 'column', padding: '2px 0px 6px 0px', background: 'rgba(0, 0, 0, 0.024)'}}>
        <div style={{display: 'flex', justifyContent: 'flex-end', gap: '4px', marginRight: '8px'}}>
        <Link href={`/user/${note.author_id}`} type="secondary">{note.author_name}</Link>
        <Text type="secondary">-</Text>
        <Text type="secondary">{formatDateTimeToYYYYMMDDHHMM(note.updated_at)}</Text>
        </div>
        <div style={{display: 'flex', justifyContent: generalized ? 'space-between' : 'flex-end', alignItems: 'center'}}>
        {generalized && (
          <LinkToDestination note={note} type={type} />
        )}
        <Link href={`/user/${note.author_id}`} type="secondary" style={{textAlign: 'right', lineHeight: 1, marginRight: '8px'}}>{note.author_email}</Link>
        </div>
        </div>
      </div>
      )
}

export default Note

interface LinkToDestinationProps {
  note: Notes
  type: string
}

const LinkToDestination = (props : LinkToDestinationProps) => {
  const { note, type } = props;

  switch (type) {
    case NOTE_TYPE.expense: 
      return (
        <Link href={`/expense/${note.expense_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Expense</Link>
      )
    case NOTE_TYPE.external_company:
      return (
        <Link href={`/external-company/${note.external_company_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Company</Link>
      )
    case NOTE_TYPE.income:
      return (
        <Link href={`/income/${note.income_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Income</Link>
      )
    case NOTE_TYPE.project:
      return (
        <Link href={`/project/${note.project_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Project</Link>
      )
    case NOTE_TYPE.sub_project:
        return (
          <Link href={`/sub-project/${note.sub_project_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Sub project</Link>
        )
    case NOTE_TYPE.product:
      return (
        <Link href={`/product/${note.product_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Product</Link>
      )
    case NOTE_TYPE.invoice:
      return (
        <Link href={`/invoice/${note.invoice_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Invoice</Link>
      )
    case NOTE_TYPE.invoice_item:
      return (
        <Link href={`/invoice-item/${note.invoice_item_id}`} type="secondary" underline style={{ lineHeight: 1, marginLeft: '8px'}}>Invoice item</Link>
      )
    default:
      return null
  }

}