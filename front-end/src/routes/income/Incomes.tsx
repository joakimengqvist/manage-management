/* eslint-disable @typescript-eslint/no-explicit-any */
import { Button, Select } from 'antd';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Incomes from '../../components/economics/incomes/incomes';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { useGetProjects } from '../../hooks';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const PrivilegeDetails = () => {
    const navigate = useNavigate();
    const projects = useGetProjects();
    const userPrivileges = useGetLoggedInUserPrivileges();
    const [project, setProject] = useState('all');

    if (!projects) {
        return null
    }

    if (!hasPrivilege(userPrivileges, PRIVILEGES.privilege_read)) return null;

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
            <Button onClick={() => navigate("/create-income")}>Create new income</Button>
        </div>
        <Incomes project={project} />
    </>)
}

export default PrivilegeDetails;