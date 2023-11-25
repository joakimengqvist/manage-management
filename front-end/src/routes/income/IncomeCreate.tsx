import { useSelector } from 'react-redux';
import CreateProjectIncome from '../../components/economics/incomes/createProjectIncome';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../types/state';

const CreateIncome: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.economics_write)) return null;
    
    return <CreateProjectIncome />
}

export default CreateIncome;