/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { EyeInvisibleOutlined, EyeTwoTone } from '@ant-design/icons';
import { Button, Input, Space, Card, Select, Typography, notification } from 'antd';
import type { SelectProps } from 'antd';
import { createUser } from '../../api/users/create'
import { State } from '../../types/state';
import { appendUser } from '../../redux/applicationDataSlice';

const { Title, Text } = Typography;

interface Privilege {
    id: string;
    name: string;
    description: string;
    created_at: string;
    updated_at: string
}

const CreateUser: React.FC = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const privileges = useSelector((state : State) => state.application.privileges);
    const projects = useSelector((state : State) => state.application.projects);

    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [privilegesOptions, setPrivilegesOptions] = useState<SelectProps['options']>([]);
    const [selectedPrivilegesOptions, setSelectedPrivilegesOptions] = useState<Array<string>>([]);
    const [projectsOptions, setProjectsOptions] = useState<SelectProps['options']>([]);
    const [selectedProjectsOptions, setSelectedProjectsOptions] = useState<Array<number>>([]);

    useEffect(() => {
        const optionsPrivileges: SelectProps['options'] = [];
        privileges.forEach((privilege : Privilege) => {
            optionsPrivileges.push({
                label: privilege.name,
                value: privilege.name,
                });
        })
        setPrivilegesOptions(optionsPrivileges);

        const optionsProjects: SelectProps['options'] = [];
        projects.forEach((project : Privilege) => {
            optionsProjects.push({
                label: project.name,
                value: project.id,
                });
        })
        setProjectsOptions(optionsProjects);
    }, [projects, privileges])

    const onHandlePrivilegeChange = (value : any) => setSelectedPrivilegesOptions(value);
    const onHandleProjectsChange = (value : any) => setSelectedProjectsOptions(value);

    const onSubmit = () => {
        createUser(userId, firstName, lastName, email, selectedPrivilegesOptions, selectedProjectsOptions, password)
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
                    message: `Created user`,
                    description: 'Succesfully created user.',
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

  return (
        <Card style={{maxWidth: '400px'}}>
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
                    onChange={onHandlePrivilegeChange}
                    options={privilegesOptions}
                />
                <Text strong>Projects</Text>
                <Select
                    mode="multiple"
                    style={{ width: '100%' }}
                    placeholder="Select projects"
                    defaultValue={[]}
                    onChange={onHandleProjectsChange}
                    options={projectsOptions}
                />
                <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                    <Button type="primary" onClick={onSubmit}>Create user</Button>
                </div>
            </Space>
        </Card>
  );
};

export default CreateUser;