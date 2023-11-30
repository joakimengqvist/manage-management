import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import CreateProject from '../../components/projects/CreateProject';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const CreateNewProject = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();

    if (!hasPrivilege(userPrivileges, PRIVILEGES.project_write)) return null;
    
    return <CreateProject />
}

export default CreateNewProject;