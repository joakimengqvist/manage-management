/* eslint-disable @typescript-eslint/no-explicit-any */
import { useDispatch, useSelector } from 'react-redux';
import { Table, Button, Popconfirm, notification, Modal, Select, Typography, Col, Row, Card } from 'antd';
import { State } from '../../types/state';
import { popProject } from '../../redux/applicationDataSlice';
import { 
    QuestionCircleOutlined, 
    DeleteOutlined, 
    ZoomInOutlined,
    SettingOutlined,
} from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus } from './../tags/ProjectStatus';
import { SubProject } from '../../types/subProject';
import { deleteSubProject } from '../../api/subProjects/delete';
import Priority from '../renderHelpers/Priority';
import EstimatedDuration from '../renderHelpers/EstimatedDuration';
import { useEffect, useState } from 'react';
import { AddProjectsSubProjectConnection } from '../../api/subProjects/specialActions/addProjectsSubProjectConnection';
import { RemoveProjectsSubProjectConnection } from '../../api/subProjects/specialActions/RemoveProjectsSubProjectConnection';

const { Text, Link } = Typography;

const columns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name'
    },
    {
        title: 'Status',
        dataIndex: 'status',
        key: 'status'
    },
    {
        title: 'Priority',
        dataIndex: 'priority',
        key: 'priority',
    },
    {
        title: 'Estimated duration',
        dataIndex: 'estimated_duration',
        key: 'estimated_duration'
    },
    {
        title: 'Projects',
        dataIndex: 'projects',
        key: 'projects'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
]

const SubProjects: React.FC = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects);
    const subProjects = useSelector((state : State) => state.application.subProjects);

    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isModalLoading, setIsModalLoading] = useState(false);
    const [modalSelectedSubProjectId, setModalSubSelectedProjectId] = useState('');
    const [modalSelectedAddProjects, setModalAddSelectedProjects] = useState([]);
    const [modalRemoveSelectedProjects, setModalRemoveSelectedProjects] = useState([]);

    useEffect(() => {
        const selectedSubProject = subProjects.find(subProject => subProject.id === modalSelectedSubProjectId);

        if (!selectedSubProject || !selectedSubProject.projects || selectedSubProject.projects.length === 0) return;

        const options = selectedSubProject.projects.map((projectId : string) => {
            return {
                    label: getProjectName(projectId),
                    value: projectId
                }
        });
         
        setModalAddSelectedProjects(options);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [modalSelectedSubProjectId]);

    const getSubProjectName = (id : string) => subProjects.find(subProject => subProject.id === id)?.name;
    const getProjectName = (id : string) => projects.find(project => project.id === id)?.name;

    const onHandleModalSelectAddProjectIds = (value : any) => setModalAddSelectedProjects(value);
    const onHandleModalSelectRemoveProjectIds = (value : any) => setModalRemoveSelectedProjects(value);

    const openModal = (subProjectId : string) => {
        setModalSubSelectedProjectId(subProjectId);
        setIsModalOpen(true);
    }

    const onModalAddProjects = () => {
        setIsModalLoading(true);
        AddProjectsSubProjectConnection(userId, modalSelectedSubProjectId, modalSelectedAddProjects).then(response => {
            if (response?.error) {
                api.error({
                    message: `Error adding project to sub project`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                });
                return;
            }
            api.info({
                message: response.message,
                placement: 'bottom',
                duration: 1.4
            });
            setIsModalOpen(false);
        }).catch(error => {
            api.error({
                message: `Error adding project to sub project`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
            setIsModalOpen(false);
        });
        setIsModalLoading(false);
    }

    const onModalRemoveProjects = () => {
        setIsModalLoading(true);
        RemoveProjectsSubProjectConnection(userId, modalSelectedSubProjectId, modalRemoveSelectedProjects).then(response => {
            if (response?.error) {
                api.error({
                    message: `Error removing project from sub project`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                });
                return;
            }
            api.info({
                message: response.message,
                placement: 'bottom',
                duration: 1.4
            });
            setIsModalOpen(false);
        }).catch(error => {
            api.error({
                message: `Error removing project from sub project`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
            setIsModalOpen(false);
        });
        setIsModalLoading(false);
    }

    const onClickdeleteProject = async (id : string) => {
        await deleteSubProject(userId, id)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Deleted sub project`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                      });
                    return
                }
                api.info({
                    message: response.message,
                    placement: 'bottom',
                    duration: 1.4
                  });
                dispatch(popProject(id))
            })
            .catch(error => {
                api.error({
                    message: `Error deleting sub project`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
    };

    const projectOptions = projects.map(project => {
        return { label: project.name, value: project.id}
      }
    );


    const modalSelectedSubProject = subProjects.find(project => project.id === modalSelectedSubProjectId);
    const projectsRemoveOptions = modalSelectedSubProject ? modalSelectedSubProject.projects.map((projectId : string) => {
        return { label: getProjectName(projectId), value: projectId }
    }) : [];

    const subProjectsData: Array<any> = subProjects && subProjects.map((subProject : SubProject) => ({                 
            name: <Link href={`/sub-project/${subProject.id}`}>{subProject.name}</Link>,
            status: <ProjectStatus status={subProject.status} />,
            priority: <Priority priority={subProject.priority} />,
            estimated_duration: <EstimatedDuration duration={subProject.estimated_duration} />,
            projects: subProject.projects && subProject.projects.map((id : string) => (<Link style={{paddingLeft: '8px', paddingRight: '8px'}} href={`/project/${id}`}>{getProjectName(id)}</Link>)), 
            operations: (
                <div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                    <Button style={{ padding: '4px' }}  type="link" onClick={() => openModal(subProject.id)}><SettingOutlined /></Button>
                    <Link style={{padding:'5px'}} href={`/sub-project/${subProject.id}`}><ZoomInOutlined /></Link>
                    {hasPrivilege(userPrivileges, PRIVILEGES.sub_project_sudo) &&
                    <Popconfirm
                        placement="top"
                        title="Are you sure?"
                        description={`Do you want to delete project ${subProject.name}`}
                        onConfirm={() => onClickdeleteProject(subProject.id)}
                        icon={<QuestionCircleOutlined style={{ color: 'red' }} />}
                        okText="Yes"
                        cancelText="No"
                    >
                        <Button danger style={{ padding: '4px' }}  type="link"><DeleteOutlined /></Button>
                    </Popconfirm>
                    }
                </div>
            )
        }));

    if (!subProjectsData) return null;


    return  <>
    {contextHolder}
        <Table size="small" bordered columns={columns} dataSource={subProjectsData} />
        <Modal
            title={`Sub project ${getSubProjectName(modalSelectedSubProjectId)} settings`}
            open={isModalOpen}
            confirmLoading={isModalLoading}
            onCancel={() => setIsModalOpen(false)}
            footer={null}
        > 
        <Row>
            <Col span={24}>
                <Card style={{marginBottom: '24px', width: '100%'}}>
                    <Text>What projects do you want to add to {getSubProjectName(modalSelectedSubProjectId)}?</Text>
                    <Select
                        style={{width: '100%', marginTop: '8px'}}
                        mode="multiple"
                        options={projectOptions}
                        onChange={onHandleModalSelectAddProjectIds}
                        value={modalSelectedAddProjects}
                    />
                    <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '16px'}}>
                        <Button onClick={onModalAddProjects} type="primary">Add Project</Button>
                    </div>
                </Card>
            </Col>
        </Row>
        <Row>
            <Col span={24}>
                <Card style={{width: '100%'}}>
                    <Text>What projects do you want to remove from {getSubProjectName(modalSelectedSubProjectId)}?</Text>
                    <Select
                        style={{width: '100%', marginTop: '8px'}}
                        mode="multiple"
                        options={projectsRemoveOptions}
                        onChange={onHandleModalSelectRemoveProjectIds}
                        value={modalRemoveSelectedProjects}
                    />
                    <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '16px'}}>
                        <Button onClick={onModalRemoveProjects} danger>Remove Projects</Button>
                    </div>
                </Card>
            </Col>
        </Row>
        <Row>
            <Col span={24}>
                <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '24px'}}>
                    <Button onClick={() => setIsModalOpen(false)}>Close</Button>
                </div>
            </Col>
        </Row>
    </Modal>
    </>
}

export default SubProjects;