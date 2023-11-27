import Project from '../../components/projects/Project';
import { useSelector } from 'react-redux';
import { State } from '../../interfaces/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

const ProjectDetails = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, PRIVILEGES.project_read)) return null;

    return <Project />
}

export default ProjectDetails;