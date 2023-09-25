/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useNavigate, useParams } from 'react-router-dom'
import { Button, Card, Space, Input, Typography, notification, Popconfirm, Divider, Select } from 'antd';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { updateProject } from '../../api/projects/update';
import { State } from '../../types/state';
import { deleteProject } from '../../api/projects/delete';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from "@ant-design/icons";
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { RenderProjectStatus } from '../tags/ProjectStatus';
import { createProjectNote } from '../../api/notes/createProjectNote';
import { getAllProjectNotesByProjectId } from '../../api/notes/getAllProjectNotesByProductId';
import { deleteProjectNoteById } from '../../api/notes/deleteProjectNoteById';

const { Text, Title } = Typography;
const { TextArea } = Input;

const Project: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useSelector((state : State) => state.user)
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects)
    const [name, setName] = useState('');
    const [projectStatus, setProjectStatus] = useState('')
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [editing, setEditing] = useState(false);
    const [projectNotes, setProjectNotes] = useState([])
    const { id } =  useParams();
    const projectId = id || ''

    useEffect(() => {
        const project = projects.find((p : any) => p.id === projectId);
        if (project) {
          try {
            setName(project.name);
            setProjectStatus(project.status)
            } catch (error : any) {
              api.error({
                message: `Error fetching privilege`,
                description: error.toString(),
                placement: "bottom",
                duration: 1.4,
              });
            }
        }
        if (projectNotes && projectNotes.length === 0 && loggedInUser?.id) {
          getAllProjectNotesByProjectId(loggedInUser.id, projectId).then(response => {
            setProjectNotes(response)
          }).catch(error => {
            console.log('error fetching project notes', error)
          })
        }
      }, [projects]);

    const onHandleNameChange = (event : any) => setName(event.target.value);
    const onHandleProjectStatusChange = (value : any) => setProjectStatus(value);
    const onHandleNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
    const onHandleNoteChange = (event : any) => setNote(event.target.value);

    const onSaveEdittedProject = () => {
        updateProject(loggedInUser.id, projectId, name, projectStatus)
        .then(response => {
          if (response?.error) {
            api.error({
                message: `Updated project failed`,
                description: response.message,
                placement: 'bottom',
                duration: 1.4
              });
              return
            }
            setEditing(false)
            api.info({
                message: `Updated project`,
                description: 'Succesfully updated project.',
                placement: 'bottom',
                duration: 1.4
            });
        })
        .catch(error => {
            api.error({
                message: `Error updating project`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    }

    const onClickdeleteProject = async () => {
        await deleteProject(loggedInUser.id, projectId)
          .then(response => {
            if (response?.error) {
              api.error({
                  message: `Deleted project failed`,
                  description: response.message,
                  placement: 'bottom',
                  duration: 1.4
                });
              return
            }
            api.info({
              message: `Deleted project`,
              description: "Succesfully deleted project.",
              placement: "bottom",
              duration: 1.2,
            });
            dispatch(popProject(projectId));
            setTimeout(() => {
              navigate("/projects");
            }, 1000);
          })
          .catch((error) => {
            api.error({
              message: `Error deleting project`,
              description: error.toString(),
              placement: "bottom",
              duration: 1.4,
            });
          });
      };

      const clearNoteFields = () => {
        setNoteTitle('');
        setNote('');
      }

      const onSubmitProjectNote = () => {
        const user = {
          id: loggedInUser.id,
          name: `${loggedInUser.firstName} ${loggedInUser.lastName}`,
          email: loggedInUser.email

        }
        createProjectNote(user, projectId, noteTitle, note).then(() => {
          api.info({
            message: `Created note`,
            description: "Succesfully created note.",
            placement: "bottom",
            duration: 1.2,
          });
        }).catch(error => {
          api.error({
            message: `Error creating note`,
            description: error.toString(),
            placement: "bottom",
            duration: 1.4,
          });
        })
      }

      const onClickdeleteNote = async (noteId : string) => {
        await deleteProjectNoteById(loggedInUser.id, noteId)
          .then(response => {
            if (response?.error) {
              api.error({
                  message: `Deleted project project note`,
                  description: response.message,
                  placement: 'bottom',
                  duration: 1.4
                });
              return
            }
            api.info({
              message: `Deleted project`,
              description: "Succesfully deleted project note.",
              placement: "bottom",
              duration: 1.2,
            });
          })
          .catch((error) => {
            api.error({
              message: `Error deleting project note`,
              description: error.toString(),
              placement: "bottom",
              duration: 1.4,
            });
          });
      };

    return (
      <div style={{display: 'flex', justifyContent: 'flex-start', gap: '20px'}}>
        <Card style={{width: '400px', height: 'fit-content'}}>
            {contextHolder}
            <Title level={4}>Project information</Title>
            <Divider style={{marginTop: '0px', marginBottom: '12px'}}/>
            <Space direction="vertical" style={{width: '100%'}}>
                <Text strong>Project name</Text>
                {editing ? (
                    <Input value={name} onChange={onHandleNameChange}/>
                ) : (
                    <Text>{name}</Text>
                )}
                <Text strong>Status</Text>
                {editing ? (
                  <Select
                    style={{width: '100%'}}
                    options={[
                        {value: 'ongoing', label: 'ongoing'},
                        {value: 'cancelled', label: 'cancelled'},
                        {value: 'completed', label: 'completed'}
                    ]}
                    onChange={onHandleProjectStatusChange}
                    value={projectStatus}
                  />
                ) : (
                    <RenderProjectStatus status={projectStatus} />
                )}
                <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px', paddingTop: '4px'}}>
                    {editing && hasPrivilege(userPrivileges, PRIVILEGES.project_sudo) && (
                        <Popconfirm
                            placement="top"
                            title="Are you sure?"
                            description={`Do you want to delete user ${name}`}
                            onConfirm={onClickdeleteProject}
                            icon={<QuestionCircleOutlined style={{ color: "red" }} />}
                            okText="Yes"
                            cancelText="No"
                        >
                            <Button danger type="link">Delete</Button>
                        </Popconfirm>
                    )}
                    <Button type={editing ? "default" : "primary"} onClick={() => setEditing(!editing)}>{editing ? 'Close' : 'Edit'}</Button>
                    {editing && (<Button type="primary" onClick={onSaveEdittedProject}>Save</Button>)}
                </div>
            </Space>
        </Card>
        <Card style={{width: '400px', height: 'fit-content'}}>
        <Title level={4}>Create note</Title>
        <Divider style={{marginTop: '0px', marginBottom: '12px'}}/>
        <Text strong>Title</Text>
        <Input value={noteTitle} onChange={onHandleNoteTitleChange} />
        <Text strong>Note</Text>
        <TextArea value={note} onChange={onHandleNoteChange}/>
        <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px', marginTop: '12px'}}>
          <Button onClick={clearNoteFields} type="link">Clear</Button>
          <Button type="primary" disabled={!note || !noteTitle} onClick={onSubmitProjectNote}>Submit</Button>
        </div>
        </Card>
        {projectNotes && (
        <Card style={{width: '400px', height: 'fit-content'}}>
        <Title level={4}>Notes</Title>
        <Divider style={{marginTop: '0px', marginBottom: '8px'}}/>
        {projectNotes.length > 0 && projectNotes.map((note : any) => (
          <div style={{width: '100%'}}>
            <div style={{display: 'flex', justifyContent: 'space-between', alignItems: 'center'}}>
              <Title level={5} style={{margin: '0px'}}>{note.title}</Title>
              <Popconfirm
                  placement="top"
                  title="Are you sure?"
                  description={`Do you want to delete note ${note.title}`}
                  onConfirm={() => onClickdeleteNote(note.id)}
                  icon={<QuestionCircleOutlined style={{ color: "red" }} />}
                  okText="Yes"
                  cancelText="No"
              >
                  <Button danger type="link">Delete note</Button>
              </Popconfirm>
            </div>
            <Text>{note.note}</Text>
            <div style={{display: 'flex', justifyContent: 'flex-end', flexDirection: 'column', marginTop: '4px'}}>
            <Text style={{textAlign: 'right', lineHeight: 1.2}}>{note.author_name}</Text>
            <Text style={{textAlign: 'right', lineHeight: 1.2}}>{note.author_email}</Text>
            <Text type="secondary" style={{textAlign: 'right'}}>{note.created_at}</Text>
            </div>
            <Divider style={{marginTop: '8px', marginBottom: '8px'}}/>
          </div>
          ))}
        </Card>
        )}
    </div>)
}

export default Project;