/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, DatePicker, Row, Typography } from 'antd';
import { Button, Input, Space, Card, notification, Select } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { createSubProject } from '../../api/subProjects/create';
import { State } from '../../types/state';
import { appendProject } from '../../redux/applicationDataSlice';

const { Title, Text } = Typography;

interface CreateSubProject {
    projectIdParam: string | undefined;
}

const CreateSubProject = (props : CreateSubProject) => {
    const { projectIdParam } = props;
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const allProjects = useSelector((state: State) => state.application.projects);
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [status, setStatus] = useState('');
    const [priority, setPriority] = useState(0);
    const [startDate, setStartDate] = useState('');
    const [dueDate, setDueDate] = useState('');
    const [estimatedDuration, setEstimatedDuration] = useState(0);
    const [projectId, setProjectId] = useState('');

    useEffect(() => {
        if (projectIdParam) setProjectId(projectIdParam);
    }, [projectIdParam])

    const onHandleNameChange = (event : any) => setName(event.target.value);
    const onHandleStatusChange = (value : any) => setStatus(value);

    const onChangeStartDate = (value : any) => {
        if (value) {
            setStartDate(value.$d)
        }
    }
    const onChangeDueDate = (value : any) => {
        if (value) {
            setDueDate(value.$d)
        }
    }

    const projectOptions = allProjects.map(project => {
        return { label: project.name, value: project.id}
      }
    );



    const onSubmit = () => {
        createSubProject(
            userId, 
            name, 
            status, 
            description, 
            priority, 
            startDate, 
            dueDate, 
            estimatedDuration, 
            projectId, 
            [],
            [],
            [],
            [],
        ).then(response => {
            if (response?.error) {
                api.error({
                    message: `Create subProject failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                    });
                return
            }
            api.info({
                message: `Created subProject`,
                description: 'Succesfully created subProject.',
                placement: 'bottom',
                duration: 1.4
                });
            dispatch(appendProject({
                id: response.data,
                name: name,
                status: status,
            }))
        })
        .catch(error => {
            api.error({
                message: `Error deleting privilege`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    }

  return (
        <Card style={{maxWidth: '800px'}}>
            {contextHolder}
            <Row>
                <Col>
                <Title level={4}>Create Project</Title>
                </Col>
            </Row>
            <Row>
                <Col span={12}>
                <Space direction="vertical" style={{width: '100%', paddingRight: '20px'}}>
                
                <Text strong>Project name</Text>
                <Input 
                    placeholder="Project name" 
                    value={name} 
                    onChange={onHandleNameChange} 
                    onBlur={onHandleNameChange}
                />
                <Text strong>Project description</Text>
                <Input
                    placeholder="Project description"
                    value={description}
                    onChange={(event) => setDescription(event.target.value)}
                    onBlur={(event) => setDescription(event.target.value)}
                />
                <Text strong>Project priority</Text>
                <Select
                    style={{width: '100%'}}
                    options={[
                        {value: 1, label: 'very-low'},
                        {value: 2, label: 'low'},
                        {value: 3, label: 'neutral'},
                        {value: 4, label: 'high'},
                        {value: 5, label: 'very-high'},
                    ]}
                    onChange={(value) => setPriority(Number(value))}
                    value={priority}
                />
                <Text strong>Project status</Text>
                <Select
                    style={{width: '100%'}}
                    options={[
                        {value: 'ongoing', label: 'ongoing'},
                        {value: 'cancelled', label: 'cancelled'},
                        {value: 'completed', label: 'completed'}
                    ]}
                    onChange={onHandleStatusChange}
                    value={status}
                />
                </Space>
                </Col>
                <Col span={12}>
                <Space direction="vertical" style={{width: '100%'}}>
                <Text strong>Project start date</Text>
                <DatePicker 
                    onChange={onChangeStartDate} 
                />
                <Text strong>Project due date</Text>
                <DatePicker 
                    onChange={onChangeDueDate} 
                />
                <Text strong>Project estimated duration</Text>
                <Select
                    style={{width: '100%'}}
                    options={[
                        {value: 1, label: 'very-low'},
                        {value: 2, label: 'low'},
                        {value: 3, label: 'neutral'},
                        {value: 4, label: 'high'},
                        {value: 5, label: 'very-high'},
                    ]}
                    onChange={(value) => setEstimatedDuration(Number(value))}
                    value={estimatedDuration}
                />
                <Text strong>Project id</Text>
                <Select
                    style={{width: '100%'}}
                    options={projectOptions}
                    disabled={Boolean(projectIdParam)}
                    onChange={(value) => setProjectId(value)}
                    value={projectId}
                />
                </Space>
                </Col>
            </Row>
    
            <Row>
                <Col span={24}>
                <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '20px'}}>
                    <Button type="primary" onClick={onSubmit}>Create sub project</Button>
                </div>
                </Col>
            </Row>
             
 
        </Card>
  );
};

export default CreateSubProject;