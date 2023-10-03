/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Typography } from 'antd';
import { Button, Input, Space, Card, notification, Select } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { createProject } from '../../api/projects/create'
import { State } from '../../types/state';
import { appendProject } from '../../redux/applicationDataSlice';
import { cardShadow } from '../../enums/styles';

const { Title, Text } = Typography;

const CreateProject: React.FC = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id)
    const [name, setName] = useState('');
    const [status, setStatus] = useState('')

    const onHandleNameChange = (event : any) => setName(event.target.value);
    const onHandleStatusChange = (value : any) => setStatus(value);

    const onSubmit = () => {
        createProject(userId, name, status)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Create project failed`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                      });
                    return
                }
                api.info({
                    message: `Created project`,
                    description: 'Succesfully created project.',
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
        <Card bordered={false} style={{maxWidth: '400px', borderRadius: 0, boxShadow: cardShadow}}>
            {contextHolder}
            <Space direction="vertical" style={{width: '100%'}}>
                <Title level={4}>Create Project</Title>
                <Text strong>Project name</Text>
                <Input 
                    placeholder="Project name" 
                    value={name} 
                    onChange={onHandleNameChange} 
                    onBlur={onHandleNameChange}
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
                <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                    <Button type="primary" onClick={onSubmit}>Create project</Button>
                </div>
            </Space>
        </Card>
  );
};

export default CreateProject;