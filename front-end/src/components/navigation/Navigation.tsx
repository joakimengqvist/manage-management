import { useMemo } from 'react';
import { useNavigate } from "react-router-dom";
import {
  FunctionOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UsergroupDeleteOutlined,
  UnlockOutlined,
  ProjectOutlined,
  FundProjectionScreenOutlined,
  ReconciliationOutlined,
  DollarOutlined,
  FundOutlined
} from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Menu } from 'antd';

type MenuItem = Required<MenuProps>['items'][number];

function getItem(
  label: React.ReactNode,
  key: React.Key,
  icon?: React.ReactNode,
  onClick?: () => void,
  children?: MenuItem[],
  type?: 'group',
): MenuItem {
  return {
    onClick,
    key,
    icon,
    children,
    label,
    type,
  } as MenuItem;
}

type NavigationProps = {
    isCollapsed: boolean,
    setCollapsed: (isCollapsed : boolean) => void,
    authenticated: boolean,
}

const Navigation: React.FC<NavigationProps> = (props) => {
    const { isCollapsed, setCollapsed } = props;
    const navigate = useNavigate();

  const itemsLoggedIn: MenuItem[] = useMemo(() => ([
    getItem(isCollapsed ? 'Expand' : 'Minimize', '0', isCollapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />, () => setCollapsed(!isCollapsed)),
    getItem('People', '1', <UsergroupDeleteOutlined />, () => {}, [
      getItem('Users', 'sub1-1', <UsergroupDeleteOutlined />, () => navigate("/users")),
      getItem('Privileges', 'sub1-2', <UnlockOutlined />, () => navigate("/privileges")),
    ]),
    getItem('Projects', '2', <ProjectOutlined />, () => {}, [
      getItem('Projects', 'sub2-1', <FundProjectionScreenOutlined />, () => navigate("/projects")),
    ]),
    getItem('Economics', '3', <DollarOutlined />, () => {}, [
      getItem('Expenses', 'sub3-1', <ReconciliationOutlined />, () => navigate("/expenses")),
      getItem('Incomes', 'sub3-2', <FundOutlined />, () => navigate("/incomes")),
  ]),
    getItem('Testing', '4', <FunctionOutlined />, () => navigate("/test-endpoints")),
    getItem('Services', '5', <FunctionOutlined />, () => navigate("/services")),
// eslint-disable-next-line react-hooks/exhaustive-deps
]), [isCollapsed])

  return (
      <Menu
        defaultSelectedKeys={['1']}
        defaultOpenKeys={['sub1']}
        mode="inline"
        theme={'dark'}
        inlineCollapsed={isCollapsed}
        items={itemsLoggedIn}
        style={{borderRight: '1px solid #d9d9d9', height: '100%'}}
        inlineIndent={10}
      />
  );
};

export default Navigation;