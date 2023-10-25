/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useState, useEffect } from 'react';
import { Route, Routes, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import './App.css';
import HeaderMenu from './components/header/HeaderMenu';
import Navigation from './components/navigation/Navigation';
import Login from './routes/login';
import UsersDetails from './routes/Users';
import User from './routes/User';
import Projects from './routes/Projects';
import Project from './routes/Project';
import Privileges from './routes/Privileges';
import Privilege from './routes/Privilege';
import { Layout } from 'antd';
import { getAllPrivileges } from './api/privileges/getAll';
import { clearData, fetchExternalCompanies, fetchPrivileges, fetchProjects, fetchSubProjects, fetchUsers, initiateApplicationData } from './redux/applicationDataSlice';
import { initiateUser } from './redux/userDataSlice'
import { getAllUsers } from './api/users/getAll';
import { getAllProjects } from './api/projects/getAll';
import { State } from './types/state';
import { hasPrivilege } from './helpers/hasPrivileges';
import { PRIVILEGES } from './enums/privileges';
import MyPage from './routes/myPage';
import ExpenseCreate from './routes/ExpenseCreate';
import Expense from './routes/Expense';
import Expenses from './routes/Expenses';
import ServiceOverview from './routes/ServiceOverview';
import Incomes from './routes/Incomes';
import IncomeCreate from './routes/IncomeCreate';
import Income from './routes/Income';
import ExternalCompanies from './routes/ExternalCompanies';
import ExternalCompanyCreate from './routes/ExternalCompanyCreate';
import ExternalCompany from './routes/ExternalCompany';
import { getAllExternalCompanies } from './api/externalCompanies/getAll';
import SubProjects from './routes/SubProjects';
import SubProjectCreate from './routes/SubProjectCreate';
import { getAllSubProjects } from './api/subProjects/getAll';
import ProjectCreate from './routes/ProjectCreate';
import SubProject from './routes/SubProject';

const { Header, Sider, Content } = Layout;

const App: React.FC = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [isInitiated, setIsInitiated] = useState(false);
  const authenticated = useSelector((state : State) => state.user.authenticated);
  const loggedInUserId = useSelector((state : State) => state.user.id);
  const userPrivileges = useSelector((state : State) => state.user.privileges)
  const [collapsed, setCollapsed] = useState(false);

  if (!authenticated) {
    dispatch(clearData())
    navigate('/login')
  }

  useEffect(() => {
    if (!isInitiated) {
      dispatch(initiateUser());
      dispatch(initiateApplicationData());
      setIsInitiated(true);
    }

    if (loggedInUserId && authenticated) {
      if (hasPrivilege(userPrivileges, PRIVILEGES.privilege_read)) {
        getAllPrivileges(loggedInUserId).then(response => dispatch(fetchPrivileges(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.user_read)) {
        getAllUsers(loggedInUserId).then(response => dispatch(fetchUsers(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.project_read)) {
        getAllProjects(loggedInUserId).then(response => dispatch(fetchProjects(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read)) {
        getAllSubProjects(loggedInUserId).then(response => dispatch(fetchSubProjects(response.data))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.external_company_read)) {
        getAllExternalCompanies(loggedInUserId).then(response => dispatch(fetchExternalCompanies(response.data))).catch(() => {})
      }
    }
}, [loggedInUserId])

  return (
    <Layout>
      {authenticated && (
        <Sider trigger={null} collapsible collapsed={collapsed}>
          <Navigation authenticated={authenticated} isCollapsed={collapsed} setCollapsed={setCollapsed} />
        </Sider>
      )}
      <Layout>
        {authenticated && (
          <Header style={{ padding: 0, height: '48px'}}>
            <HeaderMenu />
          </Header>
        )}
        <Content style={{ padding: 0, minHeight: 1200, background: 'white' }}>
          <Routes>
            <Route index element={<div>Home</div>} />

            <Route path="/my-details" element={<MyPage />} />
            <Route path="/login" element={<Login />} />

            <Route path="/user/:id" element={<User />} />
            <Route path="/users" element={<UsersDetails />} />

            <Route path="/project/:id" element={<Project />} />
            <Route path="/projects" element={<Projects />} />
            <Route path="/create-project" element={<ProjectCreate />} />

            <Route path="/sub-projects" element={<SubProjects />} />
            <Route path="/create-sub-project" element={<SubProjectCreate />} />
            <Route path="/sub-project/:id" element={<SubProject />} />

            <Route path="/privileges" element={<Privileges />} />
            <Route path="/privilege/:id" element={<Privilege />} />

            <Route path="/expenses" element={<Expenses />} />
            <Route path="/create-expense" element={<ExpenseCreate />} />
            <Route path="/expense/:id" element={<Expense />} />

            <Route path="/incomes" element={<Incomes />} />
            <Route path="/create-income" element={<IncomeCreate />} />
            <Route path="/income/:id" element={<Income />} />

            <Route path="/external-companies" element={<ExternalCompanies />} />
            <Route path="/create-external-company" element={<ExternalCompanyCreate />} />
            <Route path="/external-company/:id" element={<ExternalCompany />} />

            <Route path="/services" element={<ServiceOverview />} />

          </Routes>
        </Content>
      </Layout>
    </Layout>
  );
};

export default App;
