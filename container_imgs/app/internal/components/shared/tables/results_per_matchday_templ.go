// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package tables

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"app/internal/db"
	"strconv"
)

func ResultsPerMatchday(userResults []db.UserMatchdayResults) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"overflow-x-auto\"><table class=\"min-w-full bg-gray-800 shadow-md rounded-lg overflow-hidden\"><thead class=\"bg-gray-700\"><tr><th class=\"py-3 px-4 text-left text-xs font-medium text-gray-300 uppercase tracking-wider text-center\">Jornada</th>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, user := range getUserNames(userResults) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<th class=\"py-3 px-4 text-left text-xs font-medium text-gray-300 uppercase tracking-wider text-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(user)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/results_per_matchday.templ`, Line: 15, Col: 114}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr></thead> <tbody class=\"divide-y divide-gray-700\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, matchday := range getMatchdays(userResults) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"hover:bg-gray-700 transition-colors duration-200\"><td class=\"py-3 px-4 whitespace-nowrap text-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(matchday))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/results_per_matchday.templ`, Line: 22, Col: 81}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, user := range getUserNames(userResults) {
				var templ_7745c5c3_Var4 = []any{"py-3 px-4 whitespace-nowrap text-center", getClassDebt(userResults, user, matchday)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var4).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/results_per_matchday.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(getPointsForUserAndMatchday(userResults, user, matchday))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/results_per_matchday.templ`, Line: 24, Col: 161}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func getUserNames(userResults []db.UserMatchdayResults) []string {
	userNamesMap := make(map[string]bool)
	var userNames []string
	for _, point := range userResults {
		if !userNamesMap[point.UserName] {
			userNamesMap[point.UserName] = true
			userNames = append(userNames, point.UserName)
		}
	}
	return userNames
}

func getMatchdays(userResults []db.UserMatchdayResults) []int {
	matchdaysMap := make(map[int]bool)
	var matchdays []int
	for _, point := range userResults {
		if !matchdaysMap[point.MatchDay] {
			matchdaysMap[point.MatchDay] = true
			matchdays = append(matchdays, point.MatchDay)
		}
	}
	return matchdays
}

func getPointsForUserAndMatchday(userResults []db.UserMatchdayResults, userName string, matchday int) string {
	for _, point := range userResults {
		if point.UserName == userName && point.MatchDay == matchday {
			return strconv.Itoa(point.Points)
		}
	}
	return "0"
}

func getClassDebt(userResults []db.UserMatchdayResults, userName string, matchday int) string {
	for _, result := range userResults {
		if result.UserName == userName && result.MatchDay == matchday && result.DebtEuros > 0 {
			return "text-red-600"
		}
	}
	return ""
}

var _ = templruntime.GeneratedTemplate
