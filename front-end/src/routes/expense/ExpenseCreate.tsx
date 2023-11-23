import { useSelector } from 'react-redux';
import CreateProjectExpense from '../../components/economics/expenses/createProjectExpense';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../types/state';

const CreateExpense: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.economics_write)) return null;
    
    return <CreateProjectExpense />
}

export default CreateExpense;