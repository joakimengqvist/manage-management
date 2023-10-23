import { useSelector } from 'react-redux';
import { PRIVILEGES } from '../enums/privileges';
import { hasPrivilege } from '../helpers/hasPrivileges';
import CreateProject from '../components/projects/CreateProject';
import { State } from '../types/state';

const CreateNewProject: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.project_write)) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <CreateProject />
        </div>
    )

}

export default CreateNewProject;