/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Typography } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { Button, Input, Space, Card, notification } from 'antd';
import { createPrivilege } from '../../api/privileges/create'
import { State } from '../../interfaces/state';
import { appendPrivilege } from '../../redux/applicationDataSlice';

const { Title, Text } = Typography;
const { TextArea } = Input;

const CreatePrivilege  = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const [name, setname] = useState('');
    const [description, setDescription] = useState('');

    const onSubmit = () => {
        createPrivilege(userId, name, description)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Create privilege failed`,
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
                dispatch(appendPrivilege({
                    id: response.data,
                    name: name,
                    description: description
                }))
            })
            .catch(error => {
                api.error({
                    message: `Error creating privilege`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                  });
            })
    };

  return (
        <Card>
            {contextHolder}
            <Space direction="vertical" style={{width: '100%'}}>
                <Title level={4}>Create privilege</Title>
                <Text strong>Name</Text>
                <Input 
                    placeholder="Name" 
                    value={name} 
                    onChange={event => setname(event.target.value)} 
                    onBlur={event => setname(event.target.value)}
                />
                <Text strong>Description</Text>
                <TextArea 
                    placeholder="Description" 
                    value={description} 
                    onChange={(event : any) => setDescription(event.target.value)} 
                    onBlur={(event : any) => setDescription(event.target.value)}
                />
                <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                    <Button type="primary" onClick={onSubmit}>Create privilege</Button>
                </div>
            </Space>
        </Card>
  );
};

export default CreatePrivilege;