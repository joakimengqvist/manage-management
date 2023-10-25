/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import * as React from 'react';
import { useParams } from 'react-router-dom'
import { Button, Card, Space, Typography, notification, Col, Row } from 'antd';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { ProjectStatus } from '../tags/ProjectStatus';
import { formatDateTimeToYYYYMMDDHHMM } from '../../helpers/stringDateFormatting';
import { SubProject } from '../../types/subProject';
import { getSubProjectById } from '../../api/subProjects/getById';
import UpdateSubProject from './UpdateSubProject';
import CreateNote from '../notes/CreateNote';
import { NOTE_TYPE } from '../../enums/notes';
import NoteList from '../notes/Notes';
import { createSubProjectNote } from '../../api/notes/subProject/create';
import { getAllSubProjectNotesBySubProjectId } from '../../api/notes/subProject/getAllByProjectId';

const { Text, Title, Link } = Typography;

const tabList = [
  {
    key: 'projectInformation',
    label: 'General information',
  },
  {
    key: 'subProjectFiles',
    label: 'Files',
  },
];

const Project: React.FC = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useSelector((state : State) => state.user);
    const users = useSelector((state : State) => state.application.users);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const [subProject, setSubProject] = useState<SubProject | null>(null);
    const [editing, setEditing] = useState(false);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [subProjectNotes, setSubProjectNotes] = useState([]);
    const [activeTab, setActiveTab] = useState<string>('projectInformation');
    const { id } =  useParams();
    const subProjectId = id || '';

    const onHandleChangeActiveTab = (key: string) => setActiveTab(key);
    const onHandleNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
    const onHandleNoteChange = (event : any) => setNote(event.target.value);

    useEffect(() => {
        if (!subProject) {
            getSubProjectById(loggedInUser.id, subProjectId).then(response => {
                if (response?.error) {
                    api.error({
                        message: `Get project failed`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                    });
                }
                setSubProject(response.data);
            })
            if (subProjectNotes && subProjectNotes.length === 0 && loggedInUser?.id) {
                getAllSubProjectNotesBySubProjectId(loggedInUser.id, subProjectId).then(response => {
                  if (!response.error && response.data) {
                    setSubProjectNotes(response.data)
                  }
                }).catch(error => {
                  console.log('error fetching project notes', error)
                })
              }
        }
    }, []);

    const getUserName = (id : string) => {
        const user = users.find(user => user.id === id);
        return `${user?.first_name} ${user?.last_name}`;
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
        createSubProjectNote(user, subProjectId, noteTitle, note).then((response) => {
          api.info({
            message: response.message,
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

      const contentList: Record<string, React.ReactNode> = {
        projectInformation: (
          <div style={{padding: '24px'}}>
            {!editing && (
            <div style={{display: 'flex', justifyContent: 'space-between'}}>
                <Title level={4}>Sub project</Title>
                <Button type={editing ? "default" : "primary"} onClick={() => setEditing(!editing)}>{editing ? 'Close' : 'Edit'}</Button>
            </div>
          )}
            {subProject && (<>
            {editing ? (
                <UpdateSubProject subProject={subProject} setEditing={setEditing} />
            ) : (
          <div style={{display: 'flex', justifyContent: 'flex-start', gap: '20px'}}>
          <Space direction="vertical" style={{minWidth: '320px'}}>
              <Text strong>Project name</Text>
                  <Text>{subProject.name}</Text>
              <Text strong>Status</Text>
                  <ProjectStatus status={subProject.status} />
          </Space>
          <div style={{paddingRight: '24px'}}>
          <Text strong>Project ID</Text><br />
          <Text>{subProjectId}</Text><br />
          {hasPrivilege(userPrivileges, 'user_read') && (<>
            <Text strong>Created by</Text><br />
            <Link href={`/user/${subProject.created_by}`}>{getUserName(subProject.created_by)}</Link><br />
            <Text strong>Created at</Text><br />
            <Text>{formatDateTimeToYYYYMMDDHHMM(subProject.created_at.toString())}</Text><br />
            <Text strong>Updated by</Text><br />
            <Link href={`/user/${subProject.updated_by}`}>{getUserName(subProject.updated_by)}</Link><br />
            <Text strong>Updated at</Text><br />
            <Text>{formatDateTimeToYYYYMMDDHHMM(subProject.updated_at.toString())}</Text><br />
          </>)}
          </div>
          </div>
          )}
          </>)}
          </div>
        ),
        subProjectFiles: <p>Project files</p>,
      };

    return (
      <Row>
         {contextHolder}
      <Col span={16}>
        <Card 
          style={{ width: '98%', height: 'fit-content', padding: 0}}
          tabList={tabList}
          activeTabKey={activeTab}
          bodyStyle={{padding: '0px'}}
          onTabChange={onHandleChangeActiveTab}
          >
          {contentList[activeTab]}
        </Card>
        </Col>
        <Col span={8}>
        <Card style={{ width: '100%', height: 'fit-content'}}>
          <CreateNote
            type={NOTE_TYPE.sub_project}
            title={noteTitle}
            onTitleChange={onHandleNoteTitleChange}
            note={note}
            onNoteChange={onHandleNoteChange}
            onClearNoteFields={clearNoteFields}
            onSubmit={onSubmitProjectNote}
          />
        {hasPrivilege(userPrivileges, 'note_read') && subProjectNotes && (
          <NoteList notes={subProjectNotes} type={NOTE_TYPE.sub_project} userId={loggedInUser.id} />
        )}
        </Card>
        </Col>
      </Row>
      )
}

export default Project;