import { useSelector } from 'react-redux';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import CreateProject from '../../components/projects/CreateProject';
import { State } from '../../interfaces/state';

const CreateNewProject = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.project_write)) return null;
    
    return <CreateProject />
}

export default CreateNewProject;