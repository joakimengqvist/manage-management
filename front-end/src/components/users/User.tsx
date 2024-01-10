/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useNavigate, useParams } from "react-router-dom";
import {
  Row,
  Col,
  Button,
  Card,
  Space,
  Input,
  Select,
  Typography,
  notification,
  Popconfirm,
  Divider,
} from "antd";
import { useDispatch } from "react-redux";
import { getAllProjectNotesByUserId } from '../../api/notes/project/getAllByUserId';
import { useEffect, useState } from "react";
import { updateUser } from "../../api/users/update";
import { QuestionCircleOutlined, DeleteOutlined } from "@ant-design/icons";
import { updateUserState, popUser } from "../../redux/applicationDataSlice";
import { deleteUser } from "../../api/users/delete";
import { hasPrivilege } from "../../helpers/hasPrivileges";
import { PRIVILEGES } from "../../enums/privileges";
import { NOTE_TYPE } from "../../enums/notes";
import Notes from "../notes/Notes";
import { getAllExpenseNotesByUserId } from "../../api/notes/expense/getAllByUserId";
import { getAllIncomeNotesByUserId } from "../../api/notes/income/getAllByUserId";
import { getAllExternalCompanyNotesByUserId } from "../../api/notes/externalCompany/getAllByUserId";
import {
  ProjectOutlined,
  BuildOutlined,
  ReconciliationOutlined,
  ExceptionOutlined,
  TranslationOutlined,
  BarcodeOutlined,
  FundOutlined,
  BankOutlined
} from '@ant-design/icons';
import { getAllSubProjectNotesByUserId } from "../../api/notes/subProject/getAllByUserId";
import { PurpleTags } from "../tags/PurpleTags";
import { BlueTags } from "../tags/BlueTags";
import { ExpenseNote, ExternalCompanyNote, IncomeNote, InvoiceItemNote, InvoiceNote, ProductNote, ProjectNote, SubProjectNote, User } from "../../interfaces";
import { getUserById } from "../../api";
import { useGetLoggedInUser, useGetLoggedInUserPrivileges, useGetPrivileges, useGetProjects } from "../../hooks";
import { getAllProductNotesByUserId } from "../../api/notes/product/getAllByUserId";
import { getAllInvoiceNotesByUserId } from "../../api/notes/invoice/getAllByUserId";
import { getAllInvoiceItemNotesByUserId } from "../../api/notes/invoiceItem/getAllByUserId";

const { Text, Title } = Typography;

const userNotesTabList = [
  {
    key: 'project',
    label: <ProjectOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'subProject',
    label: <BuildOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'expense',
    label: <ReconciliationOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'income',
    label: <FundOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'companies',
    label: <BankOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'products',
    label: <BarcodeOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'invoices',
    label: <ExceptionOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'invoiceItems',
    label: <TranslationOutlined style={{paddingLeft: '12px'}} />,
  },
];

const UserDetails = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [api, contextHolder] = notification.useNotification();
  const loggedInUser = useGetLoggedInUser();
  const allProjects = useGetProjects();
  const loggedInUserPrivileges = useGetLoggedInUserPrivileges();
  const allPrivileges = useGetPrivileges();

  const [projectNotes, setProjectNotes] = useState<Array<ProjectNote>>([]);
  const [subProjectNotes, setSubProjectNotes] = useState<Array<SubProjectNote>>([]);
  const [expenseNotes, setExpenseNotes] = useState<Array<ExpenseNote>>([]);
  const [incomeNotes, setIncomeNotes] = useState<Array<IncomeNote>>([]);
  const [externalCompaniesNotes, setExternalCompaniesNotes] = useState<Array<ExternalCompanyNote>>([]);
  const [productNotes, setProductNotes] = useState<Array<ProductNote>>([]);
  const [invoiceNotes, setInvoiceNotes] = useState<Array<InvoiceNote>>([]);
  const [invoiceItemNotes, setInvoiceItemNotes] = useState<Array<InvoiceItemNote>>([]);
  const [activeNotesTab, setActiveNotesTab] = useState<string>('project');
  
  const { id } = useParams();
  const userId = id || '';
  
  const [user, setUser] = useState<User | null>(null);
  const [email, setEmail] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [privileges, setPrivileges] = useState<Array<any>>([]);
  const [projects, setProjects] = useState<Array<any>>([]);
  const [editing, setEditing] = useState(false);

  useEffect(() => {
    getUserById(loggedInUser.id, userId).then(response => {
      if (response?.error) {
        api.error({
          message: `Error fetching user`,
          description: response.error.toString(),
          placement: "bottom",
          duration: 2,
        });
      }
      if (response.data) {
        setUser(response.data);
        setEmail(response.data.email);
        setFirstName(response.data.first_name);
        setLastName(response.data.last_name);
        setPrivileges(response.data.privileges);
        setProjects(response.data.projects);
      }
    }).catch((error : any) => {
      api.error({
        message: `Error fetching user`,
        description: error.toString(),
        placement: "bottom",
        duration: 2,
      });
    });
  }, []);

  useEffect(() => {
    if (loggedInUser.id) {

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.project_read)) {
        getAllProjectNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setProjectNotes(response.data);
          } else {
            setProjectNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching project notes', error);
        });
      }

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.sub_project_read)) {
        getAllSubProjectNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setSubProjectNotes(response.data);
          } else {
            setSubProjectNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching project notes', error);
        });
      }

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.economics_read)) {
        getAllExpenseNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setExpenseNotes(response.data);
          } else {
            setExpenseNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching expense notes', error);
        });

        getAllIncomeNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setIncomeNotes(response.data);
          } else {
            setIncomeNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching income notes', error);
        });
      }

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.external_company_read)) {
        getAllExternalCompanyNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setExternalCompaniesNotes(response.data);
          } else {
            setExternalCompaniesNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching external company notes', error);
        });
      }

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.product_read)) {      
        getAllProductNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setProductNotes(response.data);
          } else {
            setProductNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching external company notes', error);
        });
      }

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.invoice_read)) {      
        getAllInvoiceNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setInvoiceNotes(response.data);
          } else {
            setInvoiceNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching external company notes', error);
        });
      }

      if (hasPrivilege(loggedInUserPrivileges, PRIVILEGES.invoice_read)) {      
        getAllInvoiceItemNotesByUserId(loggedInUser.id, userId).then(response => {
          if (response.data?.length) {
            setInvoiceItemNotes(response.data);
          } else {
            setInvoiceItemNotes([]);
          }
        }).catch((error : any) => {
          console.log('error fetching external company notes', error);
        });
      }
    }
  }, [loggedInUser, activeNotesTab]);

  const onHandleEmailChange = (event: any) => setEmail(event.target.value);
  const onHandleFirstNameChange = (event: any) => setFirstName(event.target.value);
  const onHandleLastNameChange = (event: any) => setLastName(event.target.value);
  const onHandlePrivilegeChange = (value: any) => setPrivileges(value);
  const onHandleProjectsChange = (value: any) => setProjects(value);
  const onHandleChangeActiveNotesTab = (tab: string) => setActiveNotesTab(tab);

  const onSaveEdittedUser = async () => {
    await updateUser(
      loggedInUser.id,
      userId,
      firstName,
      lastName,
      email,
      privileges,
      projects,
    )
      .then(response => {
        if (response?.error) {
          api.error({
              message: `Updated user failed`,
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
        dispatch(updateUserState(response.data));
      })
      .catch((error) => {
        api.error({
          message: `Error updating user`,
          description: error.toString(),
          placement: "bottom",
          duration: 1.4,
        });
      });
  };

  const onClickdeleteUser = async (id: string) => {
    await deleteUser(loggedInUser.id, id)
      .then(response => {
        if (response?.error) {
          api.error({
              message: `Deleted user failed`,
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
        dispatch(popUser(id));
        setTimeout(() => {
          navigate("/users");
        }, 1000);
      })
      .catch((error) => {
        api.error({
          message: `Error deleting user`,
          description: error.toString(),
          placement: "bottom",
          duration: 1.4,
        });
      });
  };

  const projectOptions = Object.keys(allProjects).map(projectId => ({ 
    label: allProjects[projectId].name, 
    value: allProjects[projectId].id
  }));

  const privilegesOptions = Object.keys(allPrivileges).map(privilegeId => ({ 
      label: allPrivileges[privilegeId].name, 
      value: allPrivileges[privilegeId].id
  }));

  const notesContentList: Record<string, React.ReactNode> = {
    project: <Notes notes={projectNotes} type={NOTE_TYPE.project} userId={loggedInUser.id} generalized /> ,
    subProject: <Notes notes={subProjectNotes} type={NOTE_TYPE.sub_project} userId={loggedInUser.id} generalized /> ,
    expense:  <Notes notes={expenseNotes} type={NOTE_TYPE.expense} userId={loggedInUser.id} generalized /> ,
    income: <Notes notes={incomeNotes} type={NOTE_TYPE.income} userId={loggedInUser.id} generalized /> ,
    companies: <Notes notes={externalCompaniesNotes} type={NOTE_TYPE.external_company} userId={loggedInUser.id} generalized />,
    products: <Notes notes={productNotes} type={NOTE_TYPE.product} userId={loggedInUser.id} generalized />,
    invoices: <Notes notes={invoiceNotes} type={NOTE_TYPE.invoice} userId={loggedInUser.id} generalized />,
    invoiceItems: <Notes notes={invoiceItemNotes} type={NOTE_TYPE.invoice_item} userId={loggedInUser.id} generalized />
  }

  if (!user) {
    return <div>Loading...</div>;
  }

  return (
    <>
      {contextHolder}
      <Row>
        <Col span={16} style={{paddingRight: '8px'}}>
          <Card>
          <div style={{display: 'flex', justifyContent: 'flex-start', gap: '20px'}}>
            <div style={{paddingRight: '16px'}}>
            <Title level={4}>User information</Title>
            <Space direction="vertical">
                    <Text strong>First name</Text>
                    <Input
                      value={firstName}
                      disabled={!editing}
                      onChange={onHandleFirstNameChange}
                      style={{ width: 210 }}
                    />
                    <Text strong>Last name</Text>
                    <Input
                      value={lastName}
                      disabled={!editing}
                      onChange={onHandleLastNameChange}
                      style={{ width: 210 }}
                    />
                  <Text strong>Email</Text>
                  <Input
                    value={email}
                    disabled={!editing}
                    onChange={onHandleEmailChange}
                    style={{ width: 320 }}
                  />
            </Space>
            </div>
            <Space direction="vertical" style={{paddingRight: '24px'}}>
              <Title level={4} style={{marginBottom: '0px'}}>User Privileges</Title>
              <Text strong>Privileges connected to the user</Text>
              <Select
                disabled={!editing}
                mode="multiple"
                style={{ width: "100%" }}
                placeholder="Select privilege"
                value={privileges}
                tagRender={BlueTags}
                onChange={onHandlePrivilegeChange}
                options={privilegesOptions}
              />
              <Title level={4} style={{marginTop: '24px', marginBottom: '0px'}}>Projects</Title>
                <Text strong>Projects connected to the user</Text>
                <Select
                  disabled={!editing}
                  mode="multiple"
                  style={{ width: "100%" }}
                  placeholder="Select projects"
                  tagRender={PurpleTags}
                  value={projects}
                  onChange={onHandleProjectsChange}
                  options={projectOptions}
                />
            </Space>
            </div>
            <Divider />
            <div
            style={{
              display: "flex",
              justifyContent: "flex-end",
              gap: "8px",
            }}
          >
            {editing && hasPrivilege(loggedInUser.privileges, PRIVILEGES.user_sudo) && (
              <Popconfirm
                placement="top"
                title="Are you sure?"
                description={`Do you want to delete user ${firstName}`}
                onConfirm={() => onClickdeleteUser(userId)}
                icon={<QuestionCircleOutlined twoToneColor="red" />}
                okText="Yes"
                cancelText="No"
              >
                <Button danger><DeleteOutlined /></Button>
              </Popconfirm>
            )}
            <div
              style={{
                display: "flex",
                justifyContent: "flex-start",
                gap: "8px",
              }}
            >
              <Button
                type={editing ? "default" : "primary"}
                onClick={() => setEditing(!editing)}
              >
                {editing ? "Close" : "Edit"}
              </Button>
              {editing && (
                <Button type="primary" onClick={onSaveEdittedUser}>
                  Save
                </Button>
              )}
            </div>
          </div>
          </Card>
        </Col>
        <Col span={8}>
        <Card 
          style={{height: 'fit-content'}}
          tabList={userNotesTabList}
          onTabChange={onHandleChangeActiveNotesTab}
          >
            {notesContentList[activeNotesTab]}
          </Card>
        </Col>
      </Row>
    </>
  );
};

export default UserDetails;
