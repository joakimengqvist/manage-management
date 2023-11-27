/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { deleteProject } from '../../api/projects/delete';
import { Table, Button, Popconfirm, notification, Card, Select, Modal, Typography, Row, Col } from 'antd';
import { State } from '../../interfaces/state';
import { Project } from '../../interfaces/project';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined, SettingOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { SubProject } from '../../interfaces/subProject';
import { useEffect, useMemo, useState } from 'react';
import Priority from '../renderHelpers/Priority';
import EstimatedDuration from '../renderHelpers/EstimatedDuration';
import { deleteSubProject } from '../../api/subProjects/delete';
import { formatDateTimeToYYYYMMDD } from '../../helpers/stringDateFormatting';
import ReactApexChart from 'react-apexcharts';
import { addSubProjectsProjectConnection } from '../../api/subProjects/specialActions/addSubProjectsProjectConnection';
import { removeSubProjectsProjectConnection } from '../../api/subProjects/specialActions/removeSubProjectsProjectsConnection';
import ProjectStatus from '../status/ProjectStatus';

const { Link, Text } = Typography;

const tabList = [
    {
      key: 'projects',
      tab: 'Projects',
    },
    {
      key: 'overview',
      tab: 'Overview',
    },
  ];

const innerTabsList = [
    {
        key: 'sub_projects',
        tab: 'Sub projects'
    },
    {
        key: 'overview',
        tab: 'Overview'
    }
];

const projectColumns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
    },
    {
        title: 'Status',
        dataIndex: 'status',
        key: 'status'
    },
    {
        title: 'Sub projects',
        dataIndex: 'sub_projects',
        key: 'sub_projects'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
];

const subProjectColumns = [
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
        key: 'priority'  
    },
    {
        title: 'Estimated duration',
        dataIndex: 'estimated_duration',
        key: 'estimated_duration'  
    },
    {
        operations: 'Operations',
        dataIndex: 'operations',
        key: 'operations'
    }
];

const Projects = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects);
    const subProjects = useSelector((state : State) => state.application.subProjects);

    const [activeTab, setActiveTab] = useState('projects');
    const [innerActiveTab, setInnerActiveTab] = useState('sub_projects');

    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isModalLoading, setIsModalLoading] = useState(false);
    const [modalSelectedProjectId, setModalSelectedProjectId] = useState('');
    const [modalSelectedAddSubProjects, setModalAddSelectedSubProjects] = useState<Array<any>>([]);
    const [modalRemoveSelectedSubProjects, setModalRemoveSelectedSubProjects] = useState([]);

    const onHandleModalSelectAddSubProjectIds = (value : any) => setModalAddSelectedSubProjects(value);
    const onHandleModalSelectRemoveSubProjectIds = (value : any) => setModalRemoveSelectedSubProjects(value);

    const openModal = (subProjectId : string) => {
        setModalSelectedProjectId(subProjectId);
        setIsModalOpen(true);
    }

    const getProjectName = (id : string) => projects.find((project : Project) => project.id === id)?.name;
    const getSubProjectName = (id : string) => subProjects.find((project : SubProject) => project.id === id)?.name;


    useEffect(() => {
        const selectedProject = projects.find(project => project.id === modalSelectedProjectId);

        if (!selectedProject || !selectedProject.sub_projects || selectedProject.sub_projects.length === 0) return;

        const options = selectedProject.sub_projects.map((subProjectId : string) => {
            return {
                    label: getSubProjectName(subProjectId),
                    value: subProjectId
                }
        });
         
        setModalAddSelectedSubProjects(options);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [modalSelectedProjectId]);

    
    const onModalAddSubProjects = () => {
        setIsModalLoading(true);
        addSubProjectsProjectConnection(userId, modalSelectedAddSubProjects, modalSelectedProjectId).then(response => {
            if (response?.error) {
                api.error({
                    message: `Error adding sub projects to project`,
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
                message: `Error adding sub projects to project`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
            setIsModalOpen(false);
        });
        setIsModalLoading(false);
    }

    const onModalRemoveSubProjects = () => {
        setIsModalLoading(true);
        removeSubProjectsProjectConnection(userId, modalRemoveSelectedSubProjects, modalSelectedProjectId).then(response => {
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
        await deleteProject(userId, id)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Deleted project project`,
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
                    message: `Error deleting project`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
    };

    const projectsData: Array<any> = projects.map((project : Project) => {
        return {      
            id: project.id,              
            name: <Link id={project.id} href={`/project/${project.id}`}>{project.name}</Link>,
            status: <ProjectStatus status={project.status} />,
            sub_projects: project.sub_projects.map((subProjectId : string) => (
                <Link style={{marginRight: '8px'}} href={`/sub-project/${subProjectId}`}>{getSubProjectName(subProjectId)}</Link>)
            ),
            operations: (<div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                <Button style={{ padding: '4px'}} type="link" onClick={() => openModal(project.id)}><SettingOutlined /></Button>
                <Link style={{padding:'5px'}} href={`/project/${project.id}`}><ZoomInOutlined /></Link>
                {hasPrivilege(userPrivileges, PRIVILEGES.project_sudo) &&
                <Popconfirm
                    placement="top"
                    title="Are you sure?"
                    description={`Do you want to delete project ${project.name}`}
                    onConfirm={() => onClickdeleteProject(project.id)}
                    icon={<QuestionCircleOutlined twoToneColor="red" />}
                    okText="Yes"
                    cancelText="No"
                >
                    <Button style={{ padding: '4px' }} danger type="link"><DeleteOutlined /></Button>
                </Popconfirm>
                }
            </div>),
        }
    })

    const expandableProps = useMemo(() => {
        const onClickdeleteSubProject = async (id : string) => {
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

        const subProjectData = (projectId : string) => {
            if (!subProjects || subProjects.length === 0) return [];

            const subProjectForProject = subProjects.filter((subProject : SubProject) => subProject.projects.includes(projectId));
            const subProjectsDataReturned = subProjectForProject.map((subProject : any) => {
                return {
                    name: <Link href={`/sub-project/${subProject.id}`}>{subProject.name}</Link>,
                    status: <ProjectStatus status={subProject.status} />,
                    priority: <Priority priority={subProject.priority} />,
                    estimated_duration: <EstimatedDuration duration={subProject.estimated_duration} />,
                    operations: (<div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                        <Link style={{padding: '5px'}} href={`/sub-project/${subProject.id}`}><ZoomInOutlined /></Link>
                        {hasPrivilege(userPrivileges, PRIVILEGES.project_sudo) &&
                            <Popconfirm
                                placement="top"
                                title="Are you sure?"
                                description={`Do you want to delete sub project ${subProject.name}`}
                                onConfirm={() => onClickdeleteSubProject(subProject.id)}
                                icon={<QuestionCircleOutlined twoToneColor="red" />}
                                okText="Yes"
                                cancelText="No"
                            >
                                <Button style={{padding: '4px'}} danger type="link"><DeleteOutlined /></Button>
                            </Popconfirm>
                        }
                    </div>),
                }
            });
            return subProjectsDataReturned
        }



        return {
            columnWidth: 48,
            expandedRowRender: (record : any) => {

                const subProjectsForProject = subProjects.filter((subProject : SubProject) => subProject.projects.includes(record.name.props.id));
                
                const data : Array<any> = [];
                subProjectsForProject.forEach((subProject : any) => {

                    data.push({
                        x: subProject.name,
                        y: [new Date(formatDateTimeToYYYYMMDD(subProject.start_date)).getTime(), new Date(formatDateTimeToYYYYMMDD(subProject.due_date)).getTime()],
                    });
                }); 
                    
                const series = [
                    {
                      name: getProjectName(record.name.props.id),
                      data: data
                    }
                  ];

                const options =  {
                    chart: {
                        height: 450,
                    },
                    plotOptions: {
                        bar: {
                            horizontal: true,
                            barHeight: '80%'
                        }
                    },
                    xAxis: {
                        type: 'datetime'
                    },
                    stroke: {
                        width: 1
                    },
                    fill: {
                        type: 'solid',
                        opacity: 0.6
                    },
                };

                const contentList: Record<string, React.ReactNode> = {
                    sub_projects: (
                        <Table 
                            size="small" 
                            style={{marginTop: '16px'}}
                            bordered 
                            columns={subProjectColumns} 
                            dataSource={subProjectData(record.name.props.id)} 
                            summary={() => (
                                <Table.Summary>
                                    <Table.Summary.Row>
                                        <Table.Summary.Cell index={1} colSpan={5}>
                                        <div style={{display: 'flex', justifyContent: 'flex-end', gap: '16px', padding: '12px 0px 8px 0px'}}>
                                            <Button onClick={() => navigate("/create-expense")}>New expense</Button>
                                            <Button onClick={() => navigate("/create-income")}>New income</Button>
                                            <Button onClick={() => navigate("/create-sub-project")}>New sub project</Button>
                                        </div>
                                        </Table.Summary.Cell>
                                    </Table.Summary.Row>
                                </Table.Summary>
                            )}
                        />
                    ),
                    overview: <ReactApexChart options={options} series={series} type="rangeBar" height={450} />
                }
            
                return (<Card
                    style={{ height: 'fit-content', padding: 0, margin: 0, borderRadius: '0px' }}
                    bodyStyle={{padding: '0px', margin: 0}}
                    activeTabKey={innerActiveTab}
                    tabList={innerTabsList}
                    onTabChange={(key) => { setInnerActiveTab(key); }}
                >
                    {contentList[innerActiveTab]}
                </Card>
            )},
        };
      // eslint-disable-next-line react-hooks/exhaustive-deps
      }, [subProjects, userId, userPrivileges, innerActiveTab]);

      const data : Array<any> = [];

    if (projects && projects.length > 0) {
        projects.forEach(project => {
            project.sub_projects.forEach((subProjectId : any) => {
                const subProject = subProjects.find((subProject : any) => subProject.id === subProjectId);
                if (subProject?.name && subProject?.start_date && subProject?.due_date) {
                    data.push({
                        x: subProject.name,
                        y: [new Date(formatDateTimeToYYYYMMDD(subProject.start_date)).getTime(), new Date(formatDateTimeToYYYYMMDD(subProject.due_date)).getTime()],
                    });
                }
            })
        });
    }
                
    const series = [
        {
            name: 'Projects',
            data: data
        }
    ];

    const options =  {
        chart: {
            height: 450,
        },
        plotOptions: {
            bar: {
                horizontal: true,
                barHeight: '80%'
            }
        },
        stroke: {
            width: 1
        },
        xAxis: {
            type: 'datetime'
        },
        fill: {
            type: 'solid',
            opacity: 0.6
        },
    };

    const contentList: Record<string, React.ReactNode> = {
        projects: <Table size="small" rowKey="id" bordered columns={projectColumns} dataSource={projectsData} expandable={expandableProps} /> ,
        overview: <div style={{marginRight: '16px'}}><ReactApexChart options={options} series={series} type="rangeBar" height={450} /></div>
    }

    const subProjectOptions = subProjects.map(subProject => {
        return { label: subProject.name, value: subProject.id}
      }
    );

    const modalSelectedProject = projects.find(project => project.id === modalSelectedProjectId);
    const subProjectsRemoveOptions = modalSelectedProject ? modalSelectedProject.sub_projects.map((subProjectId : string) => {
        return { label: getSubProjectName(subProjectId), value: subProjectId }
    }) : [];

    return (
        <Card 
        style={{ height: 'fit-content', padding: 0}}
        bodyStyle={{padding: '0px'}}
        activeTabKey={activeTab}
        tabList={tabList}
        onTabChange={(key) => { setActiveTab(key); }}
        >
             {contextHolder}
             {contentList[activeTab]}
             <Modal
            title={`Project ${getProjectName(modalSelectedProjectId)} settings`}
            open={isModalOpen}
            confirmLoading={isModalLoading}
            onCancel={() => setIsModalOpen(false)}
            footer={null}
        > 
        <Row>
            <Col span={24}>
                <Card style={{marginBottom: '24px', width: '100%'}}>
                    <Text>What sub projects do you want to add to {getProjectName(modalSelectedProjectId)}?</Text>
                    <Select
                        style={{width: '100%', marginTop: '8px'}}
                        mode="multiple"
                        options={subProjectOptions}
                        onChange={onHandleModalSelectAddSubProjectIds}
                        value={modalSelectedAddSubProjects}
                    />
                    <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '16px'}}>
                        <Button onClick={onModalAddSubProjects} type="primary">Add sub projects</Button>
                    </div>
                </Card>
            </Col>
        </Row>
        <Row>
            <Col span={24}>
                <Card style={{width: '100%'}}>
                    <Text>What sub projects do you want to remove from {getProjectName(modalSelectedProjectId)}?</Text>
                    <Select
                        style={{width: '100%', marginTop: '8px'}}
                        mode="multiple"
                        options={subProjectsRemoveOptions}
                        onChange={onHandleModalSelectRemoveSubProjectIds}
                        value={modalRemoveSelectedSubProjects}
                    />
                    <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '16px'}}>
                        <Button onClick={onModalRemoveSubProjects} danger>Remove sub projects</Button>
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
     </Card>)
}

export default Projects;