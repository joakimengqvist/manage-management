/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { EyeInvisibleOutlined, EyeTwoTone } from '@ant-design/icons';
import { Button, Input, Space, Card, Select, Typography, notification } from 'antd';
import { createUser } from '../../api/users/create'
import { appendUser } from '../../redux/applicationDataSlice';
import { BlueTags } from '../tags/BlueTags';
import { PurpleTags } from '../tags/PurpleTags';
import { useGetLoggedInUserId, useGetPrivileges, useGetProjects } from '../../hooks';

const { Title, Text } = Typography;

const CreateUser = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useGetLoggedInUserId();
    const projects = useGetProjects();
    const privileges = useGetPrivileges();

    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [selectedPrivilegesOptions, setSelectedPrivilegesOptions] = useState<Array<string>>([]);
    const [selectedProjectsOptions, setSelectedProjectsOptions] = useState<Array<string>>([]);

    const onHandlePrivilegeChange = (value : any) => setSelectedPrivilegesOptions(value);
    const onHandleProjectsChange = (value : any) => setSelectedProjectsOptions(value);

    const onSubmit = () => {
        createUser(loggedInUserId, firstName, lastName, email, selectedPrivilegesOptions, selectedProjectsOptions, password)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Created user failed`,
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
                dispatch(appendUser({
                    id: response.data,
                    first_name: firstName,
                    last_name: lastName,
                    email: email,
                    privileges: selectedPrivilegesOptions,
                    active: true,
                }))
            })
            .catch(error => {
                api.error({
                    message: `Error creating user`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
    }

    const projectOptions = Object.keys(projects).map(projectId => ({ 
        label: projects[projectId].name, 
        value: projects[projectId].id
    }));

    const privilegesOptions = Object.keys(privileges).map(privilegeId => ({ 
        label: privileges[privilegeId].name, 
        value: privileges[privilegeId].id
    }));

  return (
        <Card>
            {contextHolder}
            <Space direction="vertical" style={{width: '100%'}}>
                <Title level={4}>Create user</Title>
                <Text strong>First name</Text>
                <Input 
                    placeholder="First name" 
                    value={firstName} 
                    onChange={event => setFirstName(event.target.value)} 
                    onBlur={event => setFirstName(event.target.value)}
                />
                <Text strong>Last name</Text>
                <Input 
                    placeholder="Last name" 
                    value={lastName} 
                    onChange={event => setLastName(event.target.value)} 
                    onBlur={event => setLastName(event.target.value)}
                />
                <Text strong>Email</Text>
                <Input 
                    placeholder="Email" 
                    value={email} 
                    onChange={event => setEmail(event.target.value)} 
                    onBlur={event => setEmail(event.target.value)}
                />
                <Text strong>Password</Text>
                <Input.Password
                    placeholder="input password"
                    iconRender={(visible) => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
                    value={password} 
                    onChange={event => setPassword(event.target.value)} 
                    onBlur={event => setPassword(event.target.value)}
                />
                 <Text strong>Privileges</Text>
                <Select
                    mode="multiple"
                    style={{ width: '100%' }}
                    placeholder="Select privileges"
                    tagRender={BlueTags}
                    onChange={onHandlePrivilegeChange}
                    options={privilegesOptions}
                />
                <Text strong>Projects</Text>
                <Select
                    mode="multiple"
                    style={{ width: '100%' }}
                    placeholder="Select projects"
                    tagRender={PurpleTags}
                    defaultValue={[]}
                    onChange={onHandleProjectsChange}
                    options={projectOptions}
                />
                <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                    <Button type="primary" onClick={onSubmit}>Create user</Button>
                </div>
            </Space>
        </Card>
  );
};

export default CreateUser;