import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import SubProject from '../../components/subProjects/SubProject';

const ProjectDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read)) return null;

    return hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read) && <SubProject />
}

export default ProjectDetails;