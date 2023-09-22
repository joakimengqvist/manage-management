/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useNavigate, useParams } from 'react-router-dom'
import { Button, Card, Space, Input, Typography, notification, Popconfirm } from 'antd';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { updateProject } from '../../api/projects/update';
import { State } from '../../types/state';
import { deleteProject } from '../../api/projects/delete';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from "@ant-design/icons";
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

const { Text, Title } = Typography;

const Project: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useSelector((state : State) => state.user.id)
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects)
    const [name, setName] = useState('');
    const [editing, setEditing] = useState(false)
    const { id } =  useParams();
    const projectId = id || ''

    useEffect(() => {
        const project = projects.find((p : any) => p.id === Number(projectId));
        if (project) {
          try {
            setName(project.name);
            } catch (error : any) {
              api.error({
                message: `Error fetching privilege`,
                description: error.toString(),
                placement: "bottom",
                duration: 1.4,
              });
            }
        }
      }, [projects]);

    const onHandleNameChange = (event : any) => setName(event.target.value);

    const onSaveEdittedProject = () => {
        updateProject(loggedInUserId, Number(projectId), name)
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
        await deleteProject(loggedInUserId, Number(projectId))
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
            dispatch(popProject(Number(projectId)));
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

    return (
        <Card style={{maxWidth: '400px'}}>
            {contextHolder}
            <Title level={4}>Project information</Title>
            <Space direction="vertical" style={{width: '100%'}}>
                <Text strong>Project name</Text>
                {editing ? (
                    <Input value={name} onChange={onHandleNameChange}/>
                ) : (
                    <Text>{name}</Text>
                )}
                <div style={{display: 'flex', justifyContent: editing ? 'space-between' : 'flex-end', paddingTop: '18px'}}>
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
                            <Button danger>Delete</Button>
                        </Popconfirm>
                    )}
                    <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px'}}>
                    <Button type={editing ? "default" : "primary"} onClick={() => setEditing(!editing)}>{editing ? 'Close' : 'Edit'}</Button>
                    {editing && (<Button type="primary" onClick={onSaveEdittedProject}>Save</Button>)}
                    </div>
                </div>
            </Space>
        </Card>
    )

}

export default Project;