import { useSelector } from 'react-redux';
import { State } from '../types/state';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { PRIVILEGES } from '../enums/privileges';
import SubProject from '../components/subProjects/SubProject';

const ProjectDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <div style={{padding: '12px 8px'}}>
            {hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read) &&
                <div style={{padding: '4px'}}>
                    <SubProject />
                </div>
            }
        </div>
    )

}

export default ProjectDetails;