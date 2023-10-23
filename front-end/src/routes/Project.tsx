import Project from '../components/projects/Project';
import { useSelector } from 'react-redux';
import { State } from '../types/state';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { PRIVILEGES } from '../enums/privileges';

const ProjectDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <div style={{padding: '12px 8px'}}>
            {hasPrivilege(userPrivileges, PRIVILEGES.project_read) && <Project />}
        </div>
    )

}

export default ProjectDetails;