/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import * as React from 'react';
import { useNavigate, useParams } from 'react-router-dom'
import { Button, Card, Space, Typography, notification, Popconfirm, Divider, Col, Row } from 'antd';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { State } from '../../types/state';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined, DeleteOutlined } from "@ant-design/icons";
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus } from '../tags/ProjectStatus';
import { formatDateTimeToYYYYMMDDHHMM } from '../../helpers/stringDateFormatting';
import { SubProject } from '../../types/subProject';
import { getSubProjectById } from '../../api/subProjects/getById';
import { deleteSubProject } from '../../api/subProjects/delete';

const { Text, Link } = Typography;

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
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useSelector((state : State) => state.user);
    const users = useSelector((state : State) => state.application.users);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const [subProject, setSubProject] = useState<SubProject | null>(null);
    const [editing, setEditing] = useState(false);
    const [activeTab, setActiveTab] = useState<string>('projectInformation');
    const { id } =  useParams();
    const subProjectId = id || '';

    const onHandleChangeActiveTab = (key: string) => setActiveTab(key);

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
        }
    }, []);

    const onClickdeleteProject = async () => {
        await deleteSubProject(loggedInUser.id, subProjectId)
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
            dispatch(popProject(subProjectId));
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

      const getUserName = (id : string) => {
        const user = users.find(user => user.id === id);
        return `${user?.first_name} ${user?.last_name}`;
    };

      const contentList: Record<string, React.ReactNode> = {
        projectInformation: (
          <div style={{padding: '24px'}}>
            {subProject && (
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
          <Divider />
          <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px'}}>
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
                      <Button danger type="link"><DeleteOutlined /></Button>
                  </Popconfirm>
              )}
              <Button type={editing ? "default" : "primary"} onClick={() => setEditing(!editing)}>{editing ? 'Close' : 'Edit'}</Button>
          </div>
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
        </Card>
        </Col>
      </Row>
      )
}

export default Project;