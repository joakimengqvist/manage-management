/* eslint-disable @typescript-eslint/no-explicit-any */
import { Button, Select } from 'antd';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import ExternalCompanies from '../../components/externalCompanies/ExternalCompanies';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { useGetProjects } from '../../hooks';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const ExternalCompaniesRoute = () => {
    const navigate = useNavigate();
    const projects = useGetProjects();
    const userPrivileges = useGetLoggedInUserPrivileges();
    const [project, setProject] = useState('all');

    if (!projects) {
        return null
    }

    if (!hasPrivilege(userPrivileges, PRIVILEGES.economics_read)) return null;

    const projectOptions = [{label: 'All projects', value: 'all'}]
    Object.keys(projects).map(projectId => {
        projectOptions.push({
            value: projects[projectId].id,
            label: projects[projectId].name,
        })
    });

    const onSelectProject = (value: any) => setProject(value);
    
    return (<>
        <div style={{display: 'flex', justifyContent: 'space-between', paddingBottom: '8px'}}>
            <Select
                    defaultValue={projectOptions[0].value}
                    style={{ width: 300 }}
                    options={projectOptions}
                    onSelect={onSelectProject}
            />
            <Button onClick={() => navigate("/create-external-company")}>Create new external company</Button>
        </div>
        <ExternalCompanies project={project} />
    </>)

}

export default ExternalCompaniesRoute;