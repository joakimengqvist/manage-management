import { useSelector } from "react-redux";
import ExternalCompanyDetails from "../../components/externalCompanies/ExternalCompany";
import { PRIVILEGES } from "../../enums/privileges";
import { hasPrivilege } from "../../helpers/hasPrivileges";
import { State } from "../../interfaces/state";

const ExternalCompanyDetailsPage = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.external_company_read)) return null;
    
    return <ExternalCompanyDetails />
}

export default ExternalCompanyDetailsPage;