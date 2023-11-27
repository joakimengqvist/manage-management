import { useSelector } from 'react-redux';
import { State } from '../../interfaces/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import SubProject from '../../components/subProjects/SubProject';

const ProjectDetails = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read)) return null;

    return <SubProject />
}

export default ProjectDetails;