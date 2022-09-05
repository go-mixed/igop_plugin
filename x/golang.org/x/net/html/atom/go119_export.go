// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19

package atom

import (
	q "golang.org/x/net/html/atom"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "atom",
		Path:       "golang.org/x/net/html/atom",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Atom": reflect.TypeOf((*q.Atom)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Lookup": reflect.ValueOf(q.Lookup),
			"String": reflect.ValueOf(q.String),
		},
		TypedConsts: map[string]igop.TypedConst{
			"A":                         {reflect.TypeOf(q.A), constant.MakeInt64(int64(q.A))},
			"Abbr":                      {reflect.TypeOf(q.Abbr), constant.MakeInt64(int64(q.Abbr))},
			"Accept":                    {reflect.TypeOf(q.Accept), constant.MakeInt64(int64(q.Accept))},
			"AcceptCharset":             {reflect.TypeOf(q.AcceptCharset), constant.MakeInt64(int64(q.AcceptCharset))},
			"Accesskey":                 {reflect.TypeOf(q.Accesskey), constant.MakeInt64(int64(q.Accesskey))},
			"Acronym":                   {reflect.TypeOf(q.Acronym), constant.MakeInt64(int64(q.Acronym))},
			"Action":                    {reflect.TypeOf(q.Action), constant.MakeInt64(int64(q.Action))},
			"Address":                   {reflect.TypeOf(q.Address), constant.MakeInt64(int64(q.Address))},
			"Align":                     {reflect.TypeOf(q.Align), constant.MakeInt64(int64(q.Align))},
			"Allowfullscreen":           {reflect.TypeOf(q.Allowfullscreen), constant.MakeInt64(int64(q.Allowfullscreen))},
			"Allowpaymentrequest":       {reflect.TypeOf(q.Allowpaymentrequest), constant.MakeInt64(int64(q.Allowpaymentrequest))},
			"Allowusermedia":            {reflect.TypeOf(q.Allowusermedia), constant.MakeInt64(int64(q.Allowusermedia))},
			"Alt":                       {reflect.TypeOf(q.Alt), constant.MakeInt64(int64(q.Alt))},
			"Annotation":                {reflect.TypeOf(q.Annotation), constant.MakeInt64(int64(q.Annotation))},
			"AnnotationXml":             {reflect.TypeOf(q.AnnotationXml), constant.MakeInt64(int64(q.AnnotationXml))},
			"Applet":                    {reflect.TypeOf(q.Applet), constant.MakeInt64(int64(q.Applet))},
			"Area":                      {reflect.TypeOf(q.Area), constant.MakeInt64(int64(q.Area))},
			"Article":                   {reflect.TypeOf(q.Article), constant.MakeInt64(int64(q.Article))},
			"As":                        {reflect.TypeOf(q.As), constant.MakeInt64(int64(q.As))},
			"Aside":                     {reflect.TypeOf(q.Aside), constant.MakeInt64(int64(q.Aside))},
			"Async":                     {reflect.TypeOf(q.Async), constant.MakeInt64(int64(q.Async))},
			"Audio":                     {reflect.TypeOf(q.Audio), constant.MakeInt64(int64(q.Audio))},
			"Autocomplete":              {reflect.TypeOf(q.Autocomplete), constant.MakeInt64(int64(q.Autocomplete))},
			"Autofocus":                 {reflect.TypeOf(q.Autofocus), constant.MakeInt64(int64(q.Autofocus))},
			"Autoplay":                  {reflect.TypeOf(q.Autoplay), constant.MakeInt64(int64(q.Autoplay))},
			"B":                         {reflect.TypeOf(q.B), constant.MakeInt64(int64(q.B))},
			"Base":                      {reflect.TypeOf(q.Base), constant.MakeInt64(int64(q.Base))},
			"Basefont":                  {reflect.TypeOf(q.Basefont), constant.MakeInt64(int64(q.Basefont))},
			"Bdi":                       {reflect.TypeOf(q.Bdi), constant.MakeInt64(int64(q.Bdi))},
			"Bdo":                       {reflect.TypeOf(q.Bdo), constant.MakeInt64(int64(q.Bdo))},
			"Bgsound":                   {reflect.TypeOf(q.Bgsound), constant.MakeInt64(int64(q.Bgsound))},
			"Big":                       {reflect.TypeOf(q.Big), constant.MakeInt64(int64(q.Big))},
			"Blink":                     {reflect.TypeOf(q.Blink), constant.MakeInt64(int64(q.Blink))},
			"Blockquote":                {reflect.TypeOf(q.Blockquote), constant.MakeInt64(int64(q.Blockquote))},
			"Body":                      {reflect.TypeOf(q.Body), constant.MakeInt64(int64(q.Body))},
			"Br":                        {reflect.TypeOf(q.Br), constant.MakeInt64(int64(q.Br))},
			"Button":                    {reflect.TypeOf(q.Button), constant.MakeInt64(int64(q.Button))},
			"Canvas":                    {reflect.TypeOf(q.Canvas), constant.MakeInt64(int64(q.Canvas))},
			"Caption":                   {reflect.TypeOf(q.Caption), constant.MakeInt64(int64(q.Caption))},
			"Center":                    {reflect.TypeOf(q.Center), constant.MakeInt64(int64(q.Center))},
			"Challenge":                 {reflect.TypeOf(q.Challenge), constant.MakeInt64(int64(q.Challenge))},
			"Charset":                   {reflect.TypeOf(q.Charset), constant.MakeInt64(int64(q.Charset))},
			"Checked":                   {reflect.TypeOf(q.Checked), constant.MakeInt64(int64(q.Checked))},
			"Cite":                      {reflect.TypeOf(q.Cite), constant.MakeInt64(int64(q.Cite))},
			"Class":                     {reflect.TypeOf(q.Class), constant.MakeInt64(int64(q.Class))},
			"Code":                      {reflect.TypeOf(q.Code), constant.MakeInt64(int64(q.Code))},
			"Col":                       {reflect.TypeOf(q.Col), constant.MakeInt64(int64(q.Col))},
			"Colgroup":                  {reflect.TypeOf(q.Colgroup), constant.MakeInt64(int64(q.Colgroup))},
			"Color":                     {reflect.TypeOf(q.Color), constant.MakeInt64(int64(q.Color))},
			"Cols":                      {reflect.TypeOf(q.Cols), constant.MakeInt64(int64(q.Cols))},
			"Colspan":                   {reflect.TypeOf(q.Colspan), constant.MakeInt64(int64(q.Colspan))},
			"Command":                   {reflect.TypeOf(q.Command), constant.MakeInt64(int64(q.Command))},
			"Content":                   {reflect.TypeOf(q.Content), constant.MakeInt64(int64(q.Content))},
			"Contenteditable":           {reflect.TypeOf(q.Contenteditable), constant.MakeInt64(int64(q.Contenteditable))},
			"Contextmenu":               {reflect.TypeOf(q.Contextmenu), constant.MakeInt64(int64(q.Contextmenu))},
			"Controls":                  {reflect.TypeOf(q.Controls), constant.MakeInt64(int64(q.Controls))},
			"Coords":                    {reflect.TypeOf(q.Coords), constant.MakeInt64(int64(q.Coords))},
			"Crossorigin":               {reflect.TypeOf(q.Crossorigin), constant.MakeInt64(int64(q.Crossorigin))},
			"Data":                      {reflect.TypeOf(q.Data), constant.MakeInt64(int64(q.Data))},
			"Datalist":                  {reflect.TypeOf(q.Datalist), constant.MakeInt64(int64(q.Datalist))},
			"Datetime":                  {reflect.TypeOf(q.Datetime), constant.MakeInt64(int64(q.Datetime))},
			"Dd":                        {reflect.TypeOf(q.Dd), constant.MakeInt64(int64(q.Dd))},
			"Default":                   {reflect.TypeOf(q.Default), constant.MakeInt64(int64(q.Default))},
			"Defer":                     {reflect.TypeOf(q.Defer), constant.MakeInt64(int64(q.Defer))},
			"Del":                       {reflect.TypeOf(q.Del), constant.MakeInt64(int64(q.Del))},
			"Desc":                      {reflect.TypeOf(q.Desc), constant.MakeInt64(int64(q.Desc))},
			"Details":                   {reflect.TypeOf(q.Details), constant.MakeInt64(int64(q.Details))},
			"Dfn":                       {reflect.TypeOf(q.Dfn), constant.MakeInt64(int64(q.Dfn))},
			"Dialog":                    {reflect.TypeOf(q.Dialog), constant.MakeInt64(int64(q.Dialog))},
			"Dir":                       {reflect.TypeOf(q.Dir), constant.MakeInt64(int64(q.Dir))},
			"Dirname":                   {reflect.TypeOf(q.Dirname), constant.MakeInt64(int64(q.Dirname))},
			"Disabled":                  {reflect.TypeOf(q.Disabled), constant.MakeInt64(int64(q.Disabled))},
			"Div":                       {reflect.TypeOf(q.Div), constant.MakeInt64(int64(q.Div))},
			"Dl":                        {reflect.TypeOf(q.Dl), constant.MakeInt64(int64(q.Dl))},
			"Download":                  {reflect.TypeOf(q.Download), constant.MakeInt64(int64(q.Download))},
			"Draggable":                 {reflect.TypeOf(q.Draggable), constant.MakeInt64(int64(q.Draggable))},
			"Dropzone":                  {reflect.TypeOf(q.Dropzone), constant.MakeInt64(int64(q.Dropzone))},
			"Dt":                        {reflect.TypeOf(q.Dt), constant.MakeInt64(int64(q.Dt))},
			"Em":                        {reflect.TypeOf(q.Em), constant.MakeInt64(int64(q.Em))},
			"Embed":                     {reflect.TypeOf(q.Embed), constant.MakeInt64(int64(q.Embed))},
			"Enctype":                   {reflect.TypeOf(q.Enctype), constant.MakeInt64(int64(q.Enctype))},
			"Face":                      {reflect.TypeOf(q.Face), constant.MakeInt64(int64(q.Face))},
			"Fieldset":                  {reflect.TypeOf(q.Fieldset), constant.MakeInt64(int64(q.Fieldset))},
			"Figcaption":                {reflect.TypeOf(q.Figcaption), constant.MakeInt64(int64(q.Figcaption))},
			"Figure":                    {reflect.TypeOf(q.Figure), constant.MakeInt64(int64(q.Figure))},
			"Font":                      {reflect.TypeOf(q.Font), constant.MakeInt64(int64(q.Font))},
			"Footer":                    {reflect.TypeOf(q.Footer), constant.MakeInt64(int64(q.Footer))},
			"For":                       {reflect.TypeOf(q.For), constant.MakeInt64(int64(q.For))},
			"ForeignObject":             {reflect.TypeOf(q.ForeignObject), constant.MakeInt64(int64(q.ForeignObject))},
			"Foreignobject":             {reflect.TypeOf(q.Foreignobject), constant.MakeInt64(int64(q.Foreignobject))},
			"Form":                      {reflect.TypeOf(q.Form), constant.MakeInt64(int64(q.Form))},
			"Formaction":                {reflect.TypeOf(q.Formaction), constant.MakeInt64(int64(q.Formaction))},
			"Formenctype":               {reflect.TypeOf(q.Formenctype), constant.MakeInt64(int64(q.Formenctype))},
			"Formmethod":                {reflect.TypeOf(q.Formmethod), constant.MakeInt64(int64(q.Formmethod))},
			"Formnovalidate":            {reflect.TypeOf(q.Formnovalidate), constant.MakeInt64(int64(q.Formnovalidate))},
			"Formtarget":                {reflect.TypeOf(q.Formtarget), constant.MakeInt64(int64(q.Formtarget))},
			"Frame":                     {reflect.TypeOf(q.Frame), constant.MakeInt64(int64(q.Frame))},
			"Frameset":                  {reflect.TypeOf(q.Frameset), constant.MakeInt64(int64(q.Frameset))},
			"H1":                        {reflect.TypeOf(q.H1), constant.MakeInt64(int64(q.H1))},
			"H2":                        {reflect.TypeOf(q.H2), constant.MakeInt64(int64(q.H2))},
			"H3":                        {reflect.TypeOf(q.H3), constant.MakeInt64(int64(q.H3))},
			"H4":                        {reflect.TypeOf(q.H4), constant.MakeInt64(int64(q.H4))},
			"H5":                        {reflect.TypeOf(q.H5), constant.MakeInt64(int64(q.H5))},
			"H6":                        {reflect.TypeOf(q.H6), constant.MakeInt64(int64(q.H6))},
			"Head":                      {reflect.TypeOf(q.Head), constant.MakeInt64(int64(q.Head))},
			"Header":                    {reflect.TypeOf(q.Header), constant.MakeInt64(int64(q.Header))},
			"Headers":                   {reflect.TypeOf(q.Headers), constant.MakeInt64(int64(q.Headers))},
			"Height":                    {reflect.TypeOf(q.Height), constant.MakeInt64(int64(q.Height))},
			"Hgroup":                    {reflect.TypeOf(q.Hgroup), constant.MakeInt64(int64(q.Hgroup))},
			"Hidden":                    {reflect.TypeOf(q.Hidden), constant.MakeInt64(int64(q.Hidden))},
			"High":                      {reflect.TypeOf(q.High), constant.MakeInt64(int64(q.High))},
			"Hr":                        {reflect.TypeOf(q.Hr), constant.MakeInt64(int64(q.Hr))},
			"Href":                      {reflect.TypeOf(q.Href), constant.MakeInt64(int64(q.Href))},
			"Hreflang":                  {reflect.TypeOf(q.Hreflang), constant.MakeInt64(int64(q.Hreflang))},
			"Html":                      {reflect.TypeOf(q.Html), constant.MakeInt64(int64(q.Html))},
			"HttpEquiv":                 {reflect.TypeOf(q.HttpEquiv), constant.MakeInt64(int64(q.HttpEquiv))},
			"I":                         {reflect.TypeOf(q.I), constant.MakeInt64(int64(q.I))},
			"Icon":                      {reflect.TypeOf(q.Icon), constant.MakeInt64(int64(q.Icon))},
			"Id":                        {reflect.TypeOf(q.Id), constant.MakeInt64(int64(q.Id))},
			"Iframe":                    {reflect.TypeOf(q.Iframe), constant.MakeInt64(int64(q.Iframe))},
			"Image":                     {reflect.TypeOf(q.Image), constant.MakeInt64(int64(q.Image))},
			"Img":                       {reflect.TypeOf(q.Img), constant.MakeInt64(int64(q.Img))},
			"Input":                     {reflect.TypeOf(q.Input), constant.MakeInt64(int64(q.Input))},
			"Inputmode":                 {reflect.TypeOf(q.Inputmode), constant.MakeInt64(int64(q.Inputmode))},
			"Ins":                       {reflect.TypeOf(q.Ins), constant.MakeInt64(int64(q.Ins))},
			"Integrity":                 {reflect.TypeOf(q.Integrity), constant.MakeInt64(int64(q.Integrity))},
			"Is":                        {reflect.TypeOf(q.Is), constant.MakeInt64(int64(q.Is))},
			"Isindex":                   {reflect.TypeOf(q.Isindex), constant.MakeInt64(int64(q.Isindex))},
			"Ismap":                     {reflect.TypeOf(q.Ismap), constant.MakeInt64(int64(q.Ismap))},
			"Itemid":                    {reflect.TypeOf(q.Itemid), constant.MakeInt64(int64(q.Itemid))},
			"Itemprop":                  {reflect.TypeOf(q.Itemprop), constant.MakeInt64(int64(q.Itemprop))},
			"Itemref":                   {reflect.TypeOf(q.Itemref), constant.MakeInt64(int64(q.Itemref))},
			"Itemscope":                 {reflect.TypeOf(q.Itemscope), constant.MakeInt64(int64(q.Itemscope))},
			"Itemtype":                  {reflect.TypeOf(q.Itemtype), constant.MakeInt64(int64(q.Itemtype))},
			"Kbd":                       {reflect.TypeOf(q.Kbd), constant.MakeInt64(int64(q.Kbd))},
			"Keygen":                    {reflect.TypeOf(q.Keygen), constant.MakeInt64(int64(q.Keygen))},
			"Keytype":                   {reflect.TypeOf(q.Keytype), constant.MakeInt64(int64(q.Keytype))},
			"Kind":                      {reflect.TypeOf(q.Kind), constant.MakeInt64(int64(q.Kind))},
			"Label":                     {reflect.TypeOf(q.Label), constant.MakeInt64(int64(q.Label))},
			"Lang":                      {reflect.TypeOf(q.Lang), constant.MakeInt64(int64(q.Lang))},
			"Legend":                    {reflect.TypeOf(q.Legend), constant.MakeInt64(int64(q.Legend))},
			"Li":                        {reflect.TypeOf(q.Li), constant.MakeInt64(int64(q.Li))},
			"Link":                      {reflect.TypeOf(q.Link), constant.MakeInt64(int64(q.Link))},
			"List":                      {reflect.TypeOf(q.List), constant.MakeInt64(int64(q.List))},
			"Listing":                   {reflect.TypeOf(q.Listing), constant.MakeInt64(int64(q.Listing))},
			"Loop":                      {reflect.TypeOf(q.Loop), constant.MakeInt64(int64(q.Loop))},
			"Low":                       {reflect.TypeOf(q.Low), constant.MakeInt64(int64(q.Low))},
			"Main":                      {reflect.TypeOf(q.Main), constant.MakeInt64(int64(q.Main))},
			"Malignmark":                {reflect.TypeOf(q.Malignmark), constant.MakeInt64(int64(q.Malignmark))},
			"Manifest":                  {reflect.TypeOf(q.Manifest), constant.MakeInt64(int64(q.Manifest))},
			"Map":                       {reflect.TypeOf(q.Map), constant.MakeInt64(int64(q.Map))},
			"Mark":                      {reflect.TypeOf(q.Mark), constant.MakeInt64(int64(q.Mark))},
			"Marquee":                   {reflect.TypeOf(q.Marquee), constant.MakeInt64(int64(q.Marquee))},
			"Math":                      {reflect.TypeOf(q.Math), constant.MakeInt64(int64(q.Math))},
			"Max":                       {reflect.TypeOf(q.Max), constant.MakeInt64(int64(q.Max))},
			"Maxlength":                 {reflect.TypeOf(q.Maxlength), constant.MakeInt64(int64(q.Maxlength))},
			"Media":                     {reflect.TypeOf(q.Media), constant.MakeInt64(int64(q.Media))},
			"Mediagroup":                {reflect.TypeOf(q.Mediagroup), constant.MakeInt64(int64(q.Mediagroup))},
			"Menu":                      {reflect.TypeOf(q.Menu), constant.MakeInt64(int64(q.Menu))},
			"Menuitem":                  {reflect.TypeOf(q.Menuitem), constant.MakeInt64(int64(q.Menuitem))},
			"Meta":                      {reflect.TypeOf(q.Meta), constant.MakeInt64(int64(q.Meta))},
			"Meter":                     {reflect.TypeOf(q.Meter), constant.MakeInt64(int64(q.Meter))},
			"Method":                    {reflect.TypeOf(q.Method), constant.MakeInt64(int64(q.Method))},
			"Mglyph":                    {reflect.TypeOf(q.Mglyph), constant.MakeInt64(int64(q.Mglyph))},
			"Mi":                        {reflect.TypeOf(q.Mi), constant.MakeInt64(int64(q.Mi))},
			"Min":                       {reflect.TypeOf(q.Min), constant.MakeInt64(int64(q.Min))},
			"Minlength":                 {reflect.TypeOf(q.Minlength), constant.MakeInt64(int64(q.Minlength))},
			"Mn":                        {reflect.TypeOf(q.Mn), constant.MakeInt64(int64(q.Mn))},
			"Mo":                        {reflect.TypeOf(q.Mo), constant.MakeInt64(int64(q.Mo))},
			"Ms":                        {reflect.TypeOf(q.Ms), constant.MakeInt64(int64(q.Ms))},
			"Mtext":                     {reflect.TypeOf(q.Mtext), constant.MakeInt64(int64(q.Mtext))},
			"Multiple":                  {reflect.TypeOf(q.Multiple), constant.MakeInt64(int64(q.Multiple))},
			"Muted":                     {reflect.TypeOf(q.Muted), constant.MakeInt64(int64(q.Muted))},
			"Name":                      {reflect.TypeOf(q.Name), constant.MakeInt64(int64(q.Name))},
			"Nav":                       {reflect.TypeOf(q.Nav), constant.MakeInt64(int64(q.Nav))},
			"Nobr":                      {reflect.TypeOf(q.Nobr), constant.MakeInt64(int64(q.Nobr))},
			"Noembed":                   {reflect.TypeOf(q.Noembed), constant.MakeInt64(int64(q.Noembed))},
			"Noframes":                  {reflect.TypeOf(q.Noframes), constant.MakeInt64(int64(q.Noframes))},
			"Nomodule":                  {reflect.TypeOf(q.Nomodule), constant.MakeInt64(int64(q.Nomodule))},
			"Nonce":                     {reflect.TypeOf(q.Nonce), constant.MakeInt64(int64(q.Nonce))},
			"Noscript":                  {reflect.TypeOf(q.Noscript), constant.MakeInt64(int64(q.Noscript))},
			"Novalidate":                {reflect.TypeOf(q.Novalidate), constant.MakeInt64(int64(q.Novalidate))},
			"Object":                    {reflect.TypeOf(q.Object), constant.MakeInt64(int64(q.Object))},
			"Ol":                        {reflect.TypeOf(q.Ol), constant.MakeInt64(int64(q.Ol))},
			"Onabort":                   {reflect.TypeOf(q.Onabort), constant.MakeInt64(int64(q.Onabort))},
			"Onafterprint":              {reflect.TypeOf(q.Onafterprint), constant.MakeInt64(int64(q.Onafterprint))},
			"Onautocomplete":            {reflect.TypeOf(q.Onautocomplete), constant.MakeInt64(int64(q.Onautocomplete))},
			"Onautocompleteerror":       {reflect.TypeOf(q.Onautocompleteerror), constant.MakeInt64(int64(q.Onautocompleteerror))},
			"Onauxclick":                {reflect.TypeOf(q.Onauxclick), constant.MakeInt64(int64(q.Onauxclick))},
			"Onbeforeprint":             {reflect.TypeOf(q.Onbeforeprint), constant.MakeInt64(int64(q.Onbeforeprint))},
			"Onbeforeunload":            {reflect.TypeOf(q.Onbeforeunload), constant.MakeInt64(int64(q.Onbeforeunload))},
			"Onblur":                    {reflect.TypeOf(q.Onblur), constant.MakeInt64(int64(q.Onblur))},
			"Oncancel":                  {reflect.TypeOf(q.Oncancel), constant.MakeInt64(int64(q.Oncancel))},
			"Oncanplay":                 {reflect.TypeOf(q.Oncanplay), constant.MakeInt64(int64(q.Oncanplay))},
			"Oncanplaythrough":          {reflect.TypeOf(q.Oncanplaythrough), constant.MakeInt64(int64(q.Oncanplaythrough))},
			"Onchange":                  {reflect.TypeOf(q.Onchange), constant.MakeInt64(int64(q.Onchange))},
			"Onclick":                   {reflect.TypeOf(q.Onclick), constant.MakeInt64(int64(q.Onclick))},
			"Onclose":                   {reflect.TypeOf(q.Onclose), constant.MakeInt64(int64(q.Onclose))},
			"Oncontextmenu":             {reflect.TypeOf(q.Oncontextmenu), constant.MakeInt64(int64(q.Oncontextmenu))},
			"Oncopy":                    {reflect.TypeOf(q.Oncopy), constant.MakeInt64(int64(q.Oncopy))},
			"Oncuechange":               {reflect.TypeOf(q.Oncuechange), constant.MakeInt64(int64(q.Oncuechange))},
			"Oncut":                     {reflect.TypeOf(q.Oncut), constant.MakeInt64(int64(q.Oncut))},
			"Ondblclick":                {reflect.TypeOf(q.Ondblclick), constant.MakeInt64(int64(q.Ondblclick))},
			"Ondrag":                    {reflect.TypeOf(q.Ondrag), constant.MakeInt64(int64(q.Ondrag))},
			"Ondragend":                 {reflect.TypeOf(q.Ondragend), constant.MakeInt64(int64(q.Ondragend))},
			"Ondragenter":               {reflect.TypeOf(q.Ondragenter), constant.MakeInt64(int64(q.Ondragenter))},
			"Ondragexit":                {reflect.TypeOf(q.Ondragexit), constant.MakeInt64(int64(q.Ondragexit))},
			"Ondragleave":               {reflect.TypeOf(q.Ondragleave), constant.MakeInt64(int64(q.Ondragleave))},
			"Ondragover":                {reflect.TypeOf(q.Ondragover), constant.MakeInt64(int64(q.Ondragover))},
			"Ondragstart":               {reflect.TypeOf(q.Ondragstart), constant.MakeInt64(int64(q.Ondragstart))},
			"Ondrop":                    {reflect.TypeOf(q.Ondrop), constant.MakeInt64(int64(q.Ondrop))},
			"Ondurationchange":          {reflect.TypeOf(q.Ondurationchange), constant.MakeInt64(int64(q.Ondurationchange))},
			"Onemptied":                 {reflect.TypeOf(q.Onemptied), constant.MakeInt64(int64(q.Onemptied))},
			"Onended":                   {reflect.TypeOf(q.Onended), constant.MakeInt64(int64(q.Onended))},
			"Onerror":                   {reflect.TypeOf(q.Onerror), constant.MakeInt64(int64(q.Onerror))},
			"Onfocus":                   {reflect.TypeOf(q.Onfocus), constant.MakeInt64(int64(q.Onfocus))},
			"Onhashchange":              {reflect.TypeOf(q.Onhashchange), constant.MakeInt64(int64(q.Onhashchange))},
			"Oninput":                   {reflect.TypeOf(q.Oninput), constant.MakeInt64(int64(q.Oninput))},
			"Oninvalid":                 {reflect.TypeOf(q.Oninvalid), constant.MakeInt64(int64(q.Oninvalid))},
			"Onkeydown":                 {reflect.TypeOf(q.Onkeydown), constant.MakeInt64(int64(q.Onkeydown))},
			"Onkeypress":                {reflect.TypeOf(q.Onkeypress), constant.MakeInt64(int64(q.Onkeypress))},
			"Onkeyup":                   {reflect.TypeOf(q.Onkeyup), constant.MakeInt64(int64(q.Onkeyup))},
			"Onlanguagechange":          {reflect.TypeOf(q.Onlanguagechange), constant.MakeInt64(int64(q.Onlanguagechange))},
			"Onload":                    {reflect.TypeOf(q.Onload), constant.MakeInt64(int64(q.Onload))},
			"Onloadeddata":              {reflect.TypeOf(q.Onloadeddata), constant.MakeInt64(int64(q.Onloadeddata))},
			"Onloadedmetadata":          {reflect.TypeOf(q.Onloadedmetadata), constant.MakeInt64(int64(q.Onloadedmetadata))},
			"Onloadend":                 {reflect.TypeOf(q.Onloadend), constant.MakeInt64(int64(q.Onloadend))},
			"Onloadstart":               {reflect.TypeOf(q.Onloadstart), constant.MakeInt64(int64(q.Onloadstart))},
			"Onmessage":                 {reflect.TypeOf(q.Onmessage), constant.MakeInt64(int64(q.Onmessage))},
			"Onmessageerror":            {reflect.TypeOf(q.Onmessageerror), constant.MakeInt64(int64(q.Onmessageerror))},
			"Onmousedown":               {reflect.TypeOf(q.Onmousedown), constant.MakeInt64(int64(q.Onmousedown))},
			"Onmouseenter":              {reflect.TypeOf(q.Onmouseenter), constant.MakeInt64(int64(q.Onmouseenter))},
			"Onmouseleave":              {reflect.TypeOf(q.Onmouseleave), constant.MakeInt64(int64(q.Onmouseleave))},
			"Onmousemove":               {reflect.TypeOf(q.Onmousemove), constant.MakeInt64(int64(q.Onmousemove))},
			"Onmouseout":                {reflect.TypeOf(q.Onmouseout), constant.MakeInt64(int64(q.Onmouseout))},
			"Onmouseover":               {reflect.TypeOf(q.Onmouseover), constant.MakeInt64(int64(q.Onmouseover))},
			"Onmouseup":                 {reflect.TypeOf(q.Onmouseup), constant.MakeInt64(int64(q.Onmouseup))},
			"Onmousewheel":              {reflect.TypeOf(q.Onmousewheel), constant.MakeInt64(int64(q.Onmousewheel))},
			"Onoffline":                 {reflect.TypeOf(q.Onoffline), constant.MakeInt64(int64(q.Onoffline))},
			"Ononline":                  {reflect.TypeOf(q.Ononline), constant.MakeInt64(int64(q.Ononline))},
			"Onpagehide":                {reflect.TypeOf(q.Onpagehide), constant.MakeInt64(int64(q.Onpagehide))},
			"Onpageshow":                {reflect.TypeOf(q.Onpageshow), constant.MakeInt64(int64(q.Onpageshow))},
			"Onpaste":                   {reflect.TypeOf(q.Onpaste), constant.MakeInt64(int64(q.Onpaste))},
			"Onpause":                   {reflect.TypeOf(q.Onpause), constant.MakeInt64(int64(q.Onpause))},
			"Onplay":                    {reflect.TypeOf(q.Onplay), constant.MakeInt64(int64(q.Onplay))},
			"Onplaying":                 {reflect.TypeOf(q.Onplaying), constant.MakeInt64(int64(q.Onplaying))},
			"Onpopstate":                {reflect.TypeOf(q.Onpopstate), constant.MakeInt64(int64(q.Onpopstate))},
			"Onprogress":                {reflect.TypeOf(q.Onprogress), constant.MakeInt64(int64(q.Onprogress))},
			"Onratechange":              {reflect.TypeOf(q.Onratechange), constant.MakeInt64(int64(q.Onratechange))},
			"Onrejectionhandled":        {reflect.TypeOf(q.Onrejectionhandled), constant.MakeInt64(int64(q.Onrejectionhandled))},
			"Onreset":                   {reflect.TypeOf(q.Onreset), constant.MakeInt64(int64(q.Onreset))},
			"Onresize":                  {reflect.TypeOf(q.Onresize), constant.MakeInt64(int64(q.Onresize))},
			"Onscroll":                  {reflect.TypeOf(q.Onscroll), constant.MakeInt64(int64(q.Onscroll))},
			"Onsecuritypolicyviolation": {reflect.TypeOf(q.Onsecuritypolicyviolation), constant.MakeInt64(int64(q.Onsecuritypolicyviolation))},
			"Onseeked":                  {reflect.TypeOf(q.Onseeked), constant.MakeInt64(int64(q.Onseeked))},
			"Onseeking":                 {reflect.TypeOf(q.Onseeking), constant.MakeInt64(int64(q.Onseeking))},
			"Onselect":                  {reflect.TypeOf(q.Onselect), constant.MakeInt64(int64(q.Onselect))},
			"Onshow":                    {reflect.TypeOf(q.Onshow), constant.MakeInt64(int64(q.Onshow))},
			"Onsort":                    {reflect.TypeOf(q.Onsort), constant.MakeInt64(int64(q.Onsort))},
			"Onstalled":                 {reflect.TypeOf(q.Onstalled), constant.MakeInt64(int64(q.Onstalled))},
			"Onstorage":                 {reflect.TypeOf(q.Onstorage), constant.MakeInt64(int64(q.Onstorage))},
			"Onsubmit":                  {reflect.TypeOf(q.Onsubmit), constant.MakeInt64(int64(q.Onsubmit))},
			"Onsuspend":                 {reflect.TypeOf(q.Onsuspend), constant.MakeInt64(int64(q.Onsuspend))},
			"Ontimeupdate":              {reflect.TypeOf(q.Ontimeupdate), constant.MakeInt64(int64(q.Ontimeupdate))},
			"Ontoggle":                  {reflect.TypeOf(q.Ontoggle), constant.MakeInt64(int64(q.Ontoggle))},
			"Onunhandledrejection":      {reflect.TypeOf(q.Onunhandledrejection), constant.MakeInt64(int64(q.Onunhandledrejection))},
			"Onunload":                  {reflect.TypeOf(q.Onunload), constant.MakeInt64(int64(q.Onunload))},
			"Onvolumechange":            {reflect.TypeOf(q.Onvolumechange), constant.MakeInt64(int64(q.Onvolumechange))},
			"Onwaiting":                 {reflect.TypeOf(q.Onwaiting), constant.MakeInt64(int64(q.Onwaiting))},
			"Onwheel":                   {reflect.TypeOf(q.Onwheel), constant.MakeInt64(int64(q.Onwheel))},
			"Open":                      {reflect.TypeOf(q.Open), constant.MakeInt64(int64(q.Open))},
			"Optgroup":                  {reflect.TypeOf(q.Optgroup), constant.MakeInt64(int64(q.Optgroup))},
			"Optimum":                   {reflect.TypeOf(q.Optimum), constant.MakeInt64(int64(q.Optimum))},
			"Option":                    {reflect.TypeOf(q.Option), constant.MakeInt64(int64(q.Option))},
			"Output":                    {reflect.TypeOf(q.Output), constant.MakeInt64(int64(q.Output))},
			"P":                         {reflect.TypeOf(q.P), constant.MakeInt64(int64(q.P))},
			"Param":                     {reflect.TypeOf(q.Param), constant.MakeInt64(int64(q.Param))},
			"Pattern":                   {reflect.TypeOf(q.Pattern), constant.MakeInt64(int64(q.Pattern))},
			"Picture":                   {reflect.TypeOf(q.Picture), constant.MakeInt64(int64(q.Picture))},
			"Ping":                      {reflect.TypeOf(q.Ping), constant.MakeInt64(int64(q.Ping))},
			"Placeholder":               {reflect.TypeOf(q.Placeholder), constant.MakeInt64(int64(q.Placeholder))},
			"Plaintext":                 {reflect.TypeOf(q.Plaintext), constant.MakeInt64(int64(q.Plaintext))},
			"Playsinline":               {reflect.TypeOf(q.Playsinline), constant.MakeInt64(int64(q.Playsinline))},
			"Poster":                    {reflect.TypeOf(q.Poster), constant.MakeInt64(int64(q.Poster))},
			"Pre":                       {reflect.TypeOf(q.Pre), constant.MakeInt64(int64(q.Pre))},
			"Preload":                   {reflect.TypeOf(q.Preload), constant.MakeInt64(int64(q.Preload))},
			"Progress":                  {reflect.TypeOf(q.Progress), constant.MakeInt64(int64(q.Progress))},
			"Prompt":                    {reflect.TypeOf(q.Prompt), constant.MakeInt64(int64(q.Prompt))},
			"Public":                    {reflect.TypeOf(q.Public), constant.MakeInt64(int64(q.Public))},
			"Q":                         {reflect.TypeOf(q.Q), constant.MakeInt64(int64(q.Q))},
			"Radiogroup":                {reflect.TypeOf(q.Radiogroup), constant.MakeInt64(int64(q.Radiogroup))},
			"Rb":                        {reflect.TypeOf(q.Rb), constant.MakeInt64(int64(q.Rb))},
			"Readonly":                  {reflect.TypeOf(q.Readonly), constant.MakeInt64(int64(q.Readonly))},
			"Referrerpolicy":            {reflect.TypeOf(q.Referrerpolicy), constant.MakeInt64(int64(q.Referrerpolicy))},
			"Rel":                       {reflect.TypeOf(q.Rel), constant.MakeInt64(int64(q.Rel))},
			"Required":                  {reflect.TypeOf(q.Required), constant.MakeInt64(int64(q.Required))},
			"Reversed":                  {reflect.TypeOf(q.Reversed), constant.MakeInt64(int64(q.Reversed))},
			"Rows":                      {reflect.TypeOf(q.Rows), constant.MakeInt64(int64(q.Rows))},
			"Rowspan":                   {reflect.TypeOf(q.Rowspan), constant.MakeInt64(int64(q.Rowspan))},
			"Rp":                        {reflect.TypeOf(q.Rp), constant.MakeInt64(int64(q.Rp))},
			"Rt":                        {reflect.TypeOf(q.Rt), constant.MakeInt64(int64(q.Rt))},
			"Rtc":                       {reflect.TypeOf(q.Rtc), constant.MakeInt64(int64(q.Rtc))},
			"Ruby":                      {reflect.TypeOf(q.Ruby), constant.MakeInt64(int64(q.Ruby))},
			"S":                         {reflect.TypeOf(q.S), constant.MakeInt64(int64(q.S))},
			"Samp":                      {reflect.TypeOf(q.Samp), constant.MakeInt64(int64(q.Samp))},
			"Sandbox":                   {reflect.TypeOf(q.Sandbox), constant.MakeInt64(int64(q.Sandbox))},
			"Scope":                     {reflect.TypeOf(q.Scope), constant.MakeInt64(int64(q.Scope))},
			"Scoped":                    {reflect.TypeOf(q.Scoped), constant.MakeInt64(int64(q.Scoped))},
			"Script":                    {reflect.TypeOf(q.Script), constant.MakeInt64(int64(q.Script))},
			"Seamless":                  {reflect.TypeOf(q.Seamless), constant.MakeInt64(int64(q.Seamless))},
			"Section":                   {reflect.TypeOf(q.Section), constant.MakeInt64(int64(q.Section))},
			"Select":                    {reflect.TypeOf(q.Select), constant.MakeInt64(int64(q.Select))},
			"Selected":                  {reflect.TypeOf(q.Selected), constant.MakeInt64(int64(q.Selected))},
			"Shape":                     {reflect.TypeOf(q.Shape), constant.MakeInt64(int64(q.Shape))},
			"Size":                      {reflect.TypeOf(q.Size), constant.MakeInt64(int64(q.Size))},
			"Sizes":                     {reflect.TypeOf(q.Sizes), constant.MakeInt64(int64(q.Sizes))},
			"Slot":                      {reflect.TypeOf(q.Slot), constant.MakeInt64(int64(q.Slot))},
			"Small":                     {reflect.TypeOf(q.Small), constant.MakeInt64(int64(q.Small))},
			"Sortable":                  {reflect.TypeOf(q.Sortable), constant.MakeInt64(int64(q.Sortable))},
			"Sorted":                    {reflect.TypeOf(q.Sorted), constant.MakeInt64(int64(q.Sorted))},
			"Source":                    {reflect.TypeOf(q.Source), constant.MakeInt64(int64(q.Source))},
			"Spacer":                    {reflect.TypeOf(q.Spacer), constant.MakeInt64(int64(q.Spacer))},
			"Span":                      {reflect.TypeOf(q.Span), constant.MakeInt64(int64(q.Span))},
			"Spellcheck":                {reflect.TypeOf(q.Spellcheck), constant.MakeInt64(int64(q.Spellcheck))},
			"Src":                       {reflect.TypeOf(q.Src), constant.MakeInt64(int64(q.Src))},
			"Srcdoc":                    {reflect.TypeOf(q.Srcdoc), constant.MakeInt64(int64(q.Srcdoc))},
			"Srclang":                   {reflect.TypeOf(q.Srclang), constant.MakeInt64(int64(q.Srclang))},
			"Srcset":                    {reflect.TypeOf(q.Srcset), constant.MakeInt64(int64(q.Srcset))},
			"Start":                     {reflect.TypeOf(q.Start), constant.MakeInt64(int64(q.Start))},
			"Step":                      {reflect.TypeOf(q.Step), constant.MakeInt64(int64(q.Step))},
			"Strike":                    {reflect.TypeOf(q.Strike), constant.MakeInt64(int64(q.Strike))},
			"Strong":                    {reflect.TypeOf(q.Strong), constant.MakeInt64(int64(q.Strong))},
			"Style":                     {reflect.TypeOf(q.Style), constant.MakeInt64(int64(q.Style))},
			"Sub":                       {reflect.TypeOf(q.Sub), constant.MakeInt64(int64(q.Sub))},
			"Summary":                   {reflect.TypeOf(q.Summary), constant.MakeInt64(int64(q.Summary))},
			"Sup":                       {reflect.TypeOf(q.Sup), constant.MakeInt64(int64(q.Sup))},
			"Svg":                       {reflect.TypeOf(q.Svg), constant.MakeInt64(int64(q.Svg))},
			"System":                    {reflect.TypeOf(q.System), constant.MakeInt64(int64(q.System))},
			"Tabindex":                  {reflect.TypeOf(q.Tabindex), constant.MakeInt64(int64(q.Tabindex))},
			"Table":                     {reflect.TypeOf(q.Table), constant.MakeInt64(int64(q.Table))},
			"Target":                    {reflect.TypeOf(q.Target), constant.MakeInt64(int64(q.Target))},
			"Tbody":                     {reflect.TypeOf(q.Tbody), constant.MakeInt64(int64(q.Tbody))},
			"Td":                        {reflect.TypeOf(q.Td), constant.MakeInt64(int64(q.Td))},
			"Template":                  {reflect.TypeOf(q.Template), constant.MakeInt64(int64(q.Template))},
			"Textarea":                  {reflect.TypeOf(q.Textarea), constant.MakeInt64(int64(q.Textarea))},
			"Tfoot":                     {reflect.TypeOf(q.Tfoot), constant.MakeInt64(int64(q.Tfoot))},
			"Th":                        {reflect.TypeOf(q.Th), constant.MakeInt64(int64(q.Th))},
			"Thead":                     {reflect.TypeOf(q.Thead), constant.MakeInt64(int64(q.Thead))},
			"Time":                      {reflect.TypeOf(q.Time), constant.MakeInt64(int64(q.Time))},
			"Title":                     {reflect.TypeOf(q.Title), constant.MakeInt64(int64(q.Title))},
			"Tr":                        {reflect.TypeOf(q.Tr), constant.MakeInt64(int64(q.Tr))},
			"Track":                     {reflect.TypeOf(q.Track), constant.MakeInt64(int64(q.Track))},
			"Translate":                 {reflect.TypeOf(q.Translate), constant.MakeInt64(int64(q.Translate))},
			"Tt":                        {reflect.TypeOf(q.Tt), constant.MakeInt64(int64(q.Tt))},
			"Type":                      {reflect.TypeOf(q.Type), constant.MakeInt64(int64(q.Type))},
			"Typemustmatch":             {reflect.TypeOf(q.Typemustmatch), constant.MakeInt64(int64(q.Typemustmatch))},
			"U":                         {reflect.TypeOf(q.U), constant.MakeInt64(int64(q.U))},
			"Ul":                        {reflect.TypeOf(q.Ul), constant.MakeInt64(int64(q.Ul))},
			"Updateviacache":            {reflect.TypeOf(q.Updateviacache), constant.MakeInt64(int64(q.Updateviacache))},
			"Usemap":                    {reflect.TypeOf(q.Usemap), constant.MakeInt64(int64(q.Usemap))},
			"Value":                     {reflect.TypeOf(q.Value), constant.MakeInt64(int64(q.Value))},
			"Var":                       {reflect.TypeOf(q.Var), constant.MakeInt64(int64(q.Var))},
			"Video":                     {reflect.TypeOf(q.Video), constant.MakeInt64(int64(q.Video))},
			"Wbr":                       {reflect.TypeOf(q.Wbr), constant.MakeInt64(int64(q.Wbr))},
			"Width":                     {reflect.TypeOf(q.Width), constant.MakeInt64(int64(q.Width))},
			"Workertype":                {reflect.TypeOf(q.Workertype), constant.MakeInt64(int64(q.Workertype))},
			"Wrap":                      {reflect.TypeOf(q.Wrap), constant.MakeInt64(int64(q.Wrap))},
			"Xmp":                       {reflect.TypeOf(q.Xmp), constant.MakeInt64(int64(q.Xmp))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
