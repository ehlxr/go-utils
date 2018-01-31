package tmpl

var Controller = `
package {{.SqlData.PackageName}}.web;

import com.google.common.base.Strings;
import {{.SqlData.PackageName}}.domain.{{.SqlData.ClassName}};
import {{.SqlData.PackageName}}.common.page.Page;
import {{.SqlData.PackageName}}.common.Constants;
import {{.SqlData.PackageName}}.service.{{.SqlData.ClassName}}Service;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;
import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping(value = "/admin/{{.SqlData.ClassNameLower}}", method = {RequestMethod.GET, RequestMethod.POST})
public class {{.SqlData.ClassName}}Controller {

    private static final Logger logger = LoggerFactory.getLogger({{.SqlData.ClassName}}Controller.class);

    @Autowired
    private {{.SqlData.ClassName}}Service {{.SqlData.ClassNameLower}}Service;

    private static Integer pageSize = 15;

    @RequestMapping(value = "/edit", method = {RequestMethod.GET})
    public String edit(@RequestParam(value = "id", defaultValue = "0") Long id,
                       Model view) {
        try {
            {{.SqlData.ClassName}} {{.SqlData.ClassNameLower}} = null;
            if (id != null && id.longValue() > 0) {
                {{.SqlData.ClassNameLower}} = {{.SqlData.ClassNameLower}}Service.query{{.SqlData.ClassName}}ById(id);
            }else{
                {{.SqlData.ClassNameLower}} = new {{.SqlData.ClassName}}();
            }
            view.addAttribute("{{.SqlData.ClassNameLower}}", {{.SqlData.ClassNameLower}});

        } catch (Exception e) {
            logger.error(e.getMessage(), e);
        }
        return "admin/{{.SqlData.ClassNameLower}}/edit";
    }

    @RequestMapping(value = "/view", method = {RequestMethod.GET})
    public String view(@RequestParam(value = "id", defaultValue = "0") Long id,
                       Model view) {
        try {
            {{.SqlData.ClassName}} {{.SqlData.ClassNameLower}} = null;
            if (id != null && id.longValue() > 0) {
                {{.SqlData.ClassNameLower}} = {{.SqlData.ClassNameLower}}Service.query{{.SqlData.ClassName}}ById(id);
            }else{
                {{.SqlData.ClassNameLower}} = new {{.SqlData.ClassName}}();
            }
            view.addAttribute("{{.SqlData.ClassNameLower}}", {{.SqlData.ClassNameLower}});

        } catch (Exception e) {
            logger.error(e.getMessage(), e);
        }
        return "admin/{{.SqlData.ClassNameLower}}/view";
    }

    @RequestMapping(value = "/delete", method = {RequestMethod.DELETE})
    @ResponseBody
    public String delete(@RequestParam(value = "id", defaultValue = "0") Long id,
                         Model view) {
        try {

            long rows = {{.SqlData.ClassNameLower}}Service.delete{{.SqlData.ClassName}}(id);

        } catch (Exception e) {
            logger.error(e.getMessage(), e);
        }
        return String.format(Constants.WEB_IFRAME_SCRIPT, "删除成功！");
    }

    @RequestMapping(value = "/save", method = {RequestMethod.POST})
    @ResponseBody
    public String save({{.SqlData.ClassName}} {{.SqlData.ClassNameLower}},
                       Model view) {
        try {

            long rows = {{.SqlData.ClassNameLower}}Service.save{{.SqlData.ClassName}}({{.SqlData.ClassNameLower}});
            view.addAttribute("{{.SqlData.ClassNameLower}}", {{.SqlData.ClassNameLower}});

        } catch (Exception e) {
            logger.error(e.getMessage(), e);
        }
        return String.format(Constants.WEB_IFRAME_SCRIPT, "保存成功！");
    }

    @RequestMapping(value = "/list", method = {RequestMethod.GET, RequestMethod.POST})
    public String list(@RequestParam(value = "page", defaultValue = "0") int page,
                       @RequestParam(value = "id", required = false) Long id,
                       Model view) {
        try {
            //查询
            Map<String, Object> search = new HashMap<String, Object>();
            if (id != null) {
                search.put("id", id);
            }

            Page<{{.SqlData.ClassName}}> pageData = {{.SqlData.ClassNameLower}}Service.query{{.SqlData.ClassName}}Page(page, pageSize,search);
            //放入page对象。
            view.addAttribute("pageData", pageData);
            view.addAttribute("id", id);


        } catch (Exception e) {
            logger.error(e.getMessage(), e);
        }
        return "admin/{{.SqlData.ClassNameLower}}/list";
    }

}
`
