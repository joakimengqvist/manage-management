import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';
import Product from '../../components/products/Product';

const PrivilegeDetails = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();

    if (!hasPrivilege(userPrivileges, PRIVILEGES.product_read)) return null;

    return <Product />
}

export default PrivilegeDetails;