/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useState, useEffect } from 'react';
import { Route, Routes, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import './App.css';
import HeaderMenu from './components/header/HeaderMenu';
import Navigation from './components/navigation/Navigation';
import Login from './routes/login';
import TestingEndpoints from './routes/TestingEndpoints';
import UsersDetails from './routes/Users';
import UserDetails from './routes/User';
import ProjectsDetails from './routes/Projects';
import ProjectDetails from './routes/Project';
import PrivilegesDetails from './routes/Privileges';
import PrivilegeDetails from './routes/Privilege';
import { Layout } from 'antd';
import { getAllPrivileges } from './api/privileges/getAll';
import { clearData, fetchPrivileges, fetchProjects, fetchUsers, initiateApplicationData } from './redux/applicationDataSlice';
import {initiateUser } from './redux/userDataSlice'
import { getAllUsers } from './api/users/getAll';
import { getAllProjects } from './api/projects/getAll';
import { State } from './types/state';
import { hasPrivilege } from './helpers/hasPrivileges';
import { PRIVILEGES } from './enums/privileges';
import MyPage from './routes/myPage';
import CreateExpense from './routes/CreateExpense';
import Expense from './routes/Expense';
import Expenses from './routes/Expenses';
import ServiceOverview from './routes/ServiceOverview';
import Incomes from './routes/Incomes';
import CreateIncome from './routes/CreateIncome';
import Income from './routes/Income';

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
        getAllPrivileges(loggedInUserId).then(response => dispatch(fetchPrivileges(response))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.user_read)) {
        getAllUsers(loggedInUserId).then(response => dispatch(fetchUsers(response))).catch(() => {})
      }
      if (hasPrivilege(userPrivileges, PRIVILEGES.project_read)) {
        getAllProjects(loggedInUserId).then(response => dispatch(fetchProjects(response))).catch(() => {})
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
            <Route path="/test-endpoints" element={<TestingEndpoints />} />
            <Route path="/login" element={<Login />} />

            <Route path="/user/:id" element={<UserDetails />} />
            <Route path="/users" element={<UsersDetails />} />

            <Route path="/project/:id" element={<ProjectDetails />} />
            <Route path="/projects" element={<ProjectsDetails />} />

            <Route path="/privileges" element={<PrivilegesDetails />} />
            <Route path="/privilege/:id" element={<PrivilegeDetails />} />

            <Route path="/expenses" element={<Expenses />} />
            <Route path="/create-expense" element={<CreateExpense />} />
            <Route path="/expense/:id" element={<Expense />} />

            <Route path="/incomes" element={<Incomes />} />
            <Route path="/create-income" element={<CreateIncome />} />
            <Route path="/income/:id" element={<Income />} />

            <Route path="/services" element={<ServiceOverview />} />

          </Routes>
        </Content>
      </Layout>
    </Layout>
  );
};

export default App;
