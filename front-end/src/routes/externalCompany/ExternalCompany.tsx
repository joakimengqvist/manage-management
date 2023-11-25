import { useSelector } from "react-redux";
import ExternalCompanyDetails from "../../components/externalCompanies/externalCompany";
import { PRIVILEGES } from "../../enums/privileges";
import { hasPrivilege } from "../../helpers/hasPrivileges";
import { State } from "../../types/state";

const ExternalCompanyDetailsPage: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.external_company_read)) return null;
    
    return <ExternalCompanyDetails />
}

export default ExternalCompanyDetailsPage;