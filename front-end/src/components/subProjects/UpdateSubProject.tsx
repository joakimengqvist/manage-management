/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Button, Col, DatePicker, Popconfirm, Row, Typography } from 'antd';
import { Input, Space, notification, Select } from 'antd';
import { subProjectStatusOptions } from '../economics/options';
import { updateSubProject } from '../../api/subProjects/update';
import { QuestionCircleOutlined, DeleteOutlined } from "@ant-design/icons";
import { SubProject } from '../../interfaces/subProject';
import { deleteSubProject } from '../../api/subProjects/delete';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { useGetLoggedInUser } from '../../hooks';

const { Title, Text } = Typography;
const { TextArea } = Input;

const UpdateSubProject = ({ subProject, setEditing } : { subProject : SubProject, setEditing : (edit :boolean) => void}) => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [status, setStatus] = useState('');
    const [priority, setPriority] = useState(0);
    const [startDate, setStartDate] = useState<any>('');
    const [dueDate, setDueDate] = useState<any>('');
    const [estimatedDuration, setEstimatedDuration] = useState(0);

    useEffect(() => {
        setName(subProject.name);
        setDescription(subProject.description);
        setStatus(subProject.status);
        setPriority(subProject.priority);
        setStartDate(subProject.start_date);
        setDueDate(subProject.due_date);
        setEstimatedDuration(subProject.estimated_duration);

    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    const onHandleNameChange = (event : any) => setName(event.target.value);
    const onHandleDescriptionChange = (event : any) => setDescription(event.target.value);
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

    const onSubmit = () => {
        updateSubProject(
            loggedInUser.id, 
            subProject.id,
            name, 
            status, 
            description, 
            priority, 
            startDate, 
            dueDate, 
            estimatedDuration, 
        ).then(response => {
            if (response?.error) {
                api.error({
                    message: `updating sub project failed`,
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
        })
        .catch(error => {
            api.error({
                message: `Error updating sub project`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    }

    const onClickdeleteProject = async () => {
        await deleteSubProject(loggedInUser.id, subProject.id)
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
              message: response.message,
              placement: "bottom",
              duration: 1.2,
            });
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
        <>
            {contextHolder}
            <Row >
                <Col span={24}>
                <div style={{display: 'flex', justifyContent: 'space-between'}}>
                    <Title level={4}>Sub project</Title>
                    <div style={{display: 'flex', gap: '8px'}}>
                    {hasPrivilege(loggedInUser.privileges, PRIVILEGES.project_sudo) && (
                  <Popconfirm
                      placement="top"
                      title="Are you sure?"
                      description={`Do you want to delete user ${name}`}
                      onConfirm={onClickdeleteProject}
                      icon={<QuestionCircleOutlined twoToneColor="red" />}
                      okText="Yes"
                      cancelText="No"
                  >
                      <Button danger type="link"><DeleteOutlined /></Button>
                  </Popconfirm>
              )}
                        <Button onClick={() => setEditing(false)}>Close</Button>
                        <Button type="primary" onClick={onSubmit}>Save</Button>
                    </div>
                    </div>
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
                <TextArea
                    placeholder="Project description"
                    value={description}
                    onChange={onHandleDescriptionChange}
                    onBlur={onHandleDescriptionChange}
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
                    options={subProjectStatusOptions}
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
                </Space>
                </Col>
            </Row>
        </>
  );
};

export default UpdateSubProject;