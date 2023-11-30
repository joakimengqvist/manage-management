/* eslint-disable @typescript-eslint/no-explicit-any */
import { Button, Select } from 'antd';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Expenses from '../../components/economics/expenses/expenses';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { useGetProjects } from '../../hooks';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const ExpensesRoute = () => {
    const navigate = useNavigate();
    const projects = useGetProjects();
    const userPrivileges = useGetLoggedInUserPrivileges();
    const [project, setProject] = useState('all');

    if (!projects) {
        return null
    }

    if (!hasPrivilege(userPrivileges, PRIVILEGES.economics_read)) return null;

    const projectOptions = Object.keys(projects).map(projectId => ({ 
        label: projects[projectId].name, 
        value: projects[projectId].id
    }));

    const onSelectProject = (value: any) => setProject(value);
    
    return (<>
        <div style={{display: 'flex', justifyContent: 'space-between', paddingBottom: '8px'}}>
            <Select
                    style={{ width: 300 }}
                    options={projectOptions}
                    onSelect={onSelectProject}
            />
            <Button onClick={() => navigate("/create-expense")}>Create new expense</Button>
        </div>
        <Expenses project={project} />
    </>)

}

export default ExpensesRoute;